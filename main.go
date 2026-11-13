package main

import (
	"bufio"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"local-media-gallery/config"
	"local-media-gallery/database"
	"local-media-gallery/scripts"
	"local-media-gallery/server"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	_ "golang.org/x/image/webp"
)

// ====== 启动与环境准备相关函数 ======

// 智能托管 imagor 容器（自动建缓存目录 + 自动创建/启动容器）
func ensureImagorContainer() error {
	mediaPath := config.Config.MEDIA_HOST_PATH
	if mediaPath == "" {
		return fmt.Errorf("MEDIA_HOST_PATH 未配置")
	}

	parentDir := filepath.Dir(mediaPath)
	baseName := filepath.Base(mediaPath)
	cachePath := filepath.Join(parentDir, baseName+"_imagor_cache")

	if err := os.MkdirAll(cachePath, 0755); err != nil {
		return fmt.Errorf("自动创建缓存目录失败 (%s): %w", cachePath, err)
	}
	fmt.Printf("✅ 缓存目录已就绪: %s\n", cachePath)

	containerName := "imagor-test"
	checkCmd := exec.Command("docker", "inspect", containerName)

	if err := checkCmd.Run(); err == nil {
		fmt.Println("检测到已有的 imagor 容器，正在启动...")
		startCmd := exec.Command("docker", "start", containerName)
		if err := startCmd.Run(); err != nil {
			return fmt.Errorf("启动已有容器失败: %w", err)
		}
	} else {
		fmt.Println("未检测到 imagor 容器，正在自动创建并启动...")
		args := []string{
			"run", "--name", containerName, "-d",
			"-p", "8000:8000",
			"-v", fmt.Sprintf("%s:/mnt/media:ro", mediaPath),
			"-v", fmt.Sprintf("%s:/mnt/cache", cachePath),
			"-e", "FILE_LOADER_BASE_DIR=/mnt/media",
			"-e", "FILE_RESULT_STORAGE_BASE_DIR=/mnt/cache",
			"-e", "HTTP_LOADER_DISABLE=1",
			"-e", "IMAGOR_UNSAFE=1",
			"-e", "IMAGOR_AUTO_WEBP=1",
			"-e", "FILE_SAFE_CHARS=--",
			"shumc/imagor",
		}
		runCmd := exec.Command("docker", args...)
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		if err := runCmd.Run(); err != nil {
			return fmt.Errorf("自动创建容器失败: %w", err)
		}
	}

	fmt.Println("✅ imagor 缩略图服务已运行")
	return nil
}

// ====== 前端构建检测 ======

func checkAndBuildFrontend() error {
	frontendDir := "./web"
	distIndex := filepath.Join(frontendDir, "dist", "index.html")
	srcDir := filepath.Join(frontendDir, "src")
	packageJson := filepath.Join(frontendDir, "package.json")

	distInfo, err := os.Stat(distIndex)
	needBuild := os.IsNotExist(err)

	if !needBuild {
		distModTime := distInfo.ModTime()
		err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if !info.IsDir() && info.ModTime().After(distModTime) {
				needBuild = true
				return fmt.Errorf("stop")
			}
			return nil
		})
		if pkgInfo, err := os.Stat(packageJson); err == nil && pkgInfo.ModTime().After(distModTime) {
			needBuild = true
		}
	}

	if !needBuild {
		fmt.Println("✅ 前端代码是最新的，跳过构建。")
		return nil
	}

	fmt.Println("🔄 检测到前端代码更新，正在执行 npm run build...")
	cmd := exec.Command("cmd", "/c", "npm", "run", "build")
	cmd.Dir = frontendDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("前端构建失败: %w", err)
	}

	fmt.Println("✅ 前端构建成功！")
	return nil
}

// ====== 清理与退出相关函数 ======

func stopImagor() {
	fmt.Println("正在停止 imagor 缩略图服务...")
	_ = exec.Command("docker", "stop", "imagor-test").Run()
	fmt.Println("✅ imagor 缩略图服务已停止")
}

func waitForExit() {
	fmt.Println("\n程序即将退出，请按回车键关闭窗口...")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func cleanup() {
	fmt.Println("\n🧹 正在清理后台服务...")
	stopImagor()
}

// ====== 主函数 ======

func main() {
	defer waitForExit()

	config.LoadConfig()

	defer cleanup()

	// 2. 智能托管 imagor
	if err := ensureImagorContainer(); err != nil {
		log.Printf("❌ 致命错误: %v\n", err)
		return
	}

	// 4. 初始化数据库
	err := database.InIt()
	if err != nil {
		log.Printf("❌ 数据库初始化失败: \n %v \n", err)
		return
	}
	defer database.CloseConnection()
	log.Println("✅ Database initialized successfully.")

	// 5. 【补上】检测并构建前端
	if err := checkAndBuildFrontend(); err != nil {
		log.Printf("❌ 前端构建出错: %v\n", err)
		return
	}

	// 6. 处理用户输入
	input, err := handleInput()
	if err != nil || (input != "1" && input != "2") {
		fmt.Println("❌ 输入错误或获取输入失败")
		return
	}

	if input == "2" {
		log.Println("Start preprocessing media folders...")
		scripts.PreprocessMediaFolder()
	}

	// 7. 启动 Web 服务器
	log.Println("Start hosting web server...")
	if err := server.Init(); err != nil {
		log.Printf("❌ 服务器启动失败: \n %v \n", err)
		return
	}
}

func handleInput() (string, error) {
	fmt.Println("What do you want to do?")
	fmt.Println("1: Start web server")
	fmt.Println("2: Preprocess media folders, and start web server")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

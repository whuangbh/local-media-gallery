# 🏠 Local Media Gallery

## 🔍 Overview
🎩 **In fancy words:**

Transform any local directory into a private, self-hosted media streaming service. This project allows you to view images and videos stored on a specific device (like a PC or an External SSD) from any phone, tablet, or laptop within your Local Area Network (LAN).

😀 **In plain words:**

This project allows you to easily host a directory containing images and videos as a local server, giving you the ability to view the media stored on a specific device using any devices with a web browser in your LAN.
This project can host any directories. However it is intentionally developed to process images and videos, and will omit any non-supported file format.

> [!IMPORTANT]
> This project is optimized for web-native formats. Non-supported file types (like `.ts` or `.mkv` without web-compatible codecs) are ignored to ensure smooth playback across all devices.

## ✨ Features
* **External Storage Support:** Host directories from internal drives, external SSDs, or USB sticks.
* **Cross-Device Viewing:** Access your library via any modern web browser on your LAN.
* **Static Thumbnails:** Assign custom thumbnails for folders and videos for a polished look.
* **Docker Ready:** Get up and running in minutes without manual environment configuration.

## ✅ Supported File Formats

### 📹 Video
This project focuses on handling media files that can natively displayed without any help on desktop and mobile browsers. For example, **.ts** files which cannot be play on mobile browsers are not supported. 

1. MP4
2. M4V
3. MOV
4. WEBM

### 🖼️ Image

1. JPG/JPEG
2. PNG
3. GIF
4. WEBP


## 🛠️ How it works
1. You first need to specify a specific directory to host. The project already contains a example directory to demonstrates its capabilities.
2. In order to capture the media within the directory, you need to preprocess/index the directory first. You'll need to select option 2 when running the project, which will handle the preprocessing automatically. **The directory need to be preprocessed every time when the contents within changes. Otherwise the change in the file system will not be reflected.**  
3. Host the server, visit port 80, and enjoy!


## Screenshots
For a directory with structure like this:
```
.
└── Detailed Example/
    ├── Folder A
    ├── Folder A.png
    ├── Folder B
    ├── Folder B.png
    ├── 18173447-hd_1080_1920_60fps.mp4
    ├── 18173447-hd_1080_1920_60fps.png
    ├── pexels-alxndrlondon-32341029.jpg
    ├── pexels-hugo-guillemard-2158157486-35205659.jpg
    └── pexels-ntare-lambert-1238500624-29783124.jpg
```

This is how it looks like in a desktop browser:

<img width="1849" height="953" alt="image" src="https://github.com/user-attachments/assets/781cf3eb-8ca0-47b6-a8b8-e3856c190886" />

This is how it loks like in a mobile browser:

<img width="401" height="885" alt="image" src="https://github.com/user-attachments/assets/8d1c3225-7513-4ef3-908a-868b3ce7624c" />


## ⚡ Quick Start
1. Clone the project.
2. Run `docker-compose up --build -d` to build and run all the containers the project needs.
3. Run `docker attach local-media-gallery-server` to attach the current terminal to the project server container. The containter would log:
    ```
    2025/12/19 07:05:42 Database initialized successfully.

    What do you want to do?

    1: Start web server

    2: Preprocess media folders, and start web server
    ```
   You may not see the log in the current terminal but you can check it in the docker app.

4. If you're running the project for the first time **or** the contents of the directory have changed since last run, type `2` in the current terminal and hit enter  to preprocess/index the directory. After the preprocessing, the server will start automatically. 
5. If it is not your first time **and** the contents of the directory haven't changed, type `1` in the current terminal and hit enter. The server will start.
6. You can view the content of the directory using any device in your LAN. 

## 💡 Recommendation
If you haven't notice, the response speed of the server is not at its best when running with Docker. If you want a better experience, I would recommend installing MySQL, Go, and NodeJS onto your device and run the project locally without docker. In this way the media will load much faster in the web browser and the preprocessing will run faster without the limitation of Docker. The **windows branch** contains the localized project which I actually runs day-to-day in Windows, you can look into it if interested.

## ⚠️ Things to Consider
I recommend taking a look into the "Detailed Example" directory inside the "example_directory" directory.
1. Be aware that the filename can affect the position of the media on the website. For example, you would usually expect an order like this:
    1. 1.jpg
    2. 2.jpg
    3. 10.jpg

    However due to the difficulties in handling natural sorting. The media will have an order like this in the website:

    1. 1.jpg
    2. 10.jpg
    3. 2.jpg

    So I would recommending adding padding zeros as prefix in the filenames to avoid the incorrect order. If you're using Windows, I wll recommend looking into [Renamer](https://www.den4b.com/products/renamer) for such functions.

2. It's very inefficient to dynamically extract thumbnails for folders and videos. Therefore this project only support setting thumbnails in a static way. Be reminded that the thumbnails for folders and videos have to be placed under the same directory and have the same name. For example, a folder named "Folder A" will have a thumbnail named "Folder A.jpg" in the same directory.

3. If you use a directory from an external storage device, be aware that Windows may perform anti-malware scan on the directory on every preprocessing (It may depends on whether or not you unplug the external device before, I'm not sure). The scan can tremendously increase the time taken for the preprocess operation. You can avoid such situation by adding the directory as an exclusion to Windows's scan, but please do it at your own cost. 

## 📅 TODO
1. Add support to 'favourites'

## 🧠 Random Thoughts
I originally developed this project without using docker, and I used it daily without docker. I added docker because I want others can quickly try out the project without dealing with the environment. If you encounter any problems when using the project with docker, I am sorry. Leave an issue and I'll see what can I do.

I used it daily with a directory on my external SSD. The project allows me to carry my media files(SSD) everywhere (traveling/aboard) while still being able to browse it's content in LAN.



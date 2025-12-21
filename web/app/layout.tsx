import type { Metadata } from "next";
import "@/app/style/reset.css";
import "@/app/globals.css";

export const metadata: Metadata = {
	title: "Local Media Gallery",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="en" data-bs-theme="dark">
			<head>
				<link
					href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.8/dist/css/bootstrap.min.css"
					rel="stylesheet"
					integrity="sha384-sRIl4kxILFvY47J16cr9ZwB07vP4J8+LH7qKQnuqkuIAvNWLzeN8tE5YBujZqJLB"
					crossOrigin="anonymous"
				/>
				<link
					rel="stylesheet"
					href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.13.1/font/bootstrap-icons.min.css"
				></link>
				<script
					src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.8/dist/js/bootstrap.bundle.min.js"
					integrity="sha384-FKyoEForCGlyvwx9Hj09JcYn3nv7wiPVlz7YYwJrWVcXK/BmnVDxM+D2scQbITxI"
					crossOrigin="anonymous"
				></script>
			</head>
			<body>{children}</body>
		</html>
	);
}

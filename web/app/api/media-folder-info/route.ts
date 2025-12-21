import { NextRequest, NextResponse } from "next/server";

export async function POST(request: NextRequest) {
	try {
		const body = await request.json();
		const { url } = body;

		let backendAddress =
			"http://localhost:" + process.env.NEXT_PUBLIC_BACKEND_PORT;

		const backendRes = await fetch(
			backendAddress + "/api/media-folder-info",
			{
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify({ url }),
			}
		);

		const data = await backendRes.json();
		return NextResponse.json(data);
	} catch (error) {
		console.dir(error);
		return NextResponse.json(
			{ success: false, error: error },
			{ status: 500 }
		);
	}
}

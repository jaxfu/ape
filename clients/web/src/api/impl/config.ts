export const BASE_URL: string = import.meta.env.PROD
  ? "/api"
  : "http://localhost:5000/api";

export async function pingServer(): Promise<boolean> {
  try {
    const res = await fetch(`${BASE_URL}/health`, {
      method: "GET",
    });

    return res.status === 200;
  } catch (e) {
    return false;
  }
}

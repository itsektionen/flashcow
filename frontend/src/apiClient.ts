export async function apiTest(): Promise<number> {
  const response = await fetch("/api/test");
  return (await response.json()).n;
}

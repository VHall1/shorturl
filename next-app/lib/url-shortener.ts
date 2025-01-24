const SERVICE_URL = "http://shortener";
type HTTPMethods = "GET" | "POST";

function makeRequest(
  path: string,
  {
    headers,
    method,
    body,
  }: {
    headers?: HeadersInit | undefined;
    body?: BodyInit | null | undefined;
    method: HTTPMethods;
  }
) {
  const requestInit: RequestInit = {
    method, // TODO: cache?
    cache: "no-cache",
  };

  if (headers) {
    requestInit.headers = headers;
  }

  if (body) {
    requestInit.body = body;
  }

  return fetch(`${SERVICE_URL}${path}`, requestInit);
}

async function getShortUrl(longUrl: string): Promise<string> {
  const response = await makeRequest("/", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ url: longUrl }),
  });

  type JSONResponse = {
    url?: string;
    error?: string;
  };
  const { url, error }: JSONResponse = await response.json();

  if (!response.ok) {
    return Promise.reject(new Error(error ?? "unhandled api error"));
  }

  if (!url) {
    return Promise.reject(new Error("api did not return a url"));
  }

  return url;
}

export { getShortUrl };

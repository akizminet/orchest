/* eslint-disable @typescript-eslint/no-explicit-any */
export class FetchError<B = any> extends Error {
  readonly status?: number;
  readonly body?: B;

  constructor(message: string, status?: number, body?: B) {
    super(message);
    this.status = status;
    this.body = body;
  }
}

const getFullUrl = (url: string) => `${__BASE_URL__}${url}`;

export const fetcher = async <T, E = any>(
  url: string,
  params?: RequestInit
) => {
  const targetUrl = getFullUrl(url);

  const response = await window.fetch(targetUrl, params);
  const responseAsString = await response.text();
  const jsonResponse =
    responseAsString === "" ? {} : JSON.parse(responseAsString);

  if (!response.ok || response.status >= 299) {
    const { message } = jsonResponse;

    return Promise.reject(
      new FetchError<E>(
        message || response.statusText,
        response.status,
        jsonResponse
      )
    );
  }

  return jsonResponse as Promise<T>;
};

export type Fetcher<T = void> = (
  url: string,
  params?: RequestInit | undefined
) => Promise<T>;

export type ContentType =
  | "application/json"
  | "application/x-www-form-urlencoded";

export const HEADER = {
  JSON: { "Content-Type": "application/json; charset=UTF-8" },
  FORM: {
    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
  },
};

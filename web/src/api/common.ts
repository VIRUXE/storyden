import { isNil, omitBy } from "lodash/fp";

export type Options = {
  url: string;
  method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  headers?: Record<string, string>;
  params?: Record<string, string | string[] | boolean>;
  data?: unknown;
  responseType?: string;
  cookie?: string;
};

export const buildPayload = (data: unknown) => {
  if (!data) {
    return undefined;
  }

  if (data instanceof File || data instanceof Blob) {
    return data;
  }

  return JSON.stringify(data);
};

const removeEmpty = omitBy(isNil);

export const cleanQuery = (
  params?: Record<string, string | string[] | boolean>,
): string => {
  if (!params) return "";

  const clean = removeEmpty(params);

  const format = new URLSearchParams(clean).toString();

  return `?${format}`;
};

export function shouldLog(status: number) {
  if (status < 400) {
    return false;
  }

  if (status === 404) {
    return false;
  }

  return true;
}

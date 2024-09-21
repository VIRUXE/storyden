/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type {
  AssetGetOKResponse,
  AssetUploadBody,
  GetInfoOKResponse,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * The version number includes the date and time of the release build as
well as a short representation of the Git commit hash.

 * @summary Get the software version string.
 */
export type getVersionResponse = {
  data: string;
  status: number;
};

export const getGetVersionUrl = () => {
  return `/version`;
};

export const getVersion = async (
  options?: RequestInit,
): Promise<getVersionResponse> => {
  return fetcher<Promise<getVersionResponse>>(getGetVersionUrl(), {
    ...options,
    method: "GET",
  });
};

/**
 * Note: the generator creates a `map[string]interface{}` if this is set to
`application/json`... so I'm just using plain text for now.

 * @summary Get the OpenAPI 3.0 specification as JSON.
 */
export type getSpecResponse = {
  data: string;
  status: number;
};

export const getGetSpecUrl = () => {
  return `/openapi.json`;
};

export const getSpec = async (
  options?: RequestInit,
): Promise<getSpecResponse> => {
  return fetcher<Promise<getSpecResponse>>(getGetSpecUrl(), {
    ...options,
    method: "GET",
  });
};

/**
 * Get the basic forum installation info such as title, description, etc.

 */
export type getInfoResponse = {
  data: GetInfoOKResponse;
  status: number;
};

export const getGetInfoUrl = () => {
  return `/info`;
};

export const getInfo = async (
  options?: RequestInit,
): Promise<getInfoResponse> => {
  return fetcher<Promise<getInfoResponse>>(getGetInfoUrl(), {
    ...options,
    method: "GET",
  });
};

/**
 * Get the logo icon image.
 */
export type iconGetResponse = {
  data: AssetGetOKResponse;
  status: number;
};

export const getIconGetUrl = (
  iconSize: "512x512" | "32x32" | "180x180" | "120x120" | "167x167" | "152x152",
) => {
  return `/info/icon/${iconSize}`;
};

export const iconGet = async (
  iconSize: "512x512" | "32x32" | "180x180" | "120x120" | "167x167" | "152x152",
  options?: RequestInit,
): Promise<iconGetResponse> => {
  return fetcher<Promise<iconGetResponse>>(getIconGetUrl(iconSize), {
    ...options,
    method: "GET",
  });
};

/**
 * Upload and process the installation's logo image.
 */
export type iconUploadResponse = {
  data: void;
  status: number;
};

export const getIconUploadUrl = () => {
  return `/info/icon`;
};

export const iconUpload = async (
  assetUploadBody: AssetUploadBody,
  options?: RequestInit,
): Promise<iconUploadResponse> => {
  return fetcher<Promise<iconUploadResponse>>(getIconUploadUrl(), {
    ...options,
    method: "POST",
    body: JSON.stringify(assetUploadBody),
  });
};
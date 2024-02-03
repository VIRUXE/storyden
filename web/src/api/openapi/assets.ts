/**
 * Generated by orval v6.24.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { Arguments, Key, SWRConfiguration } from "swr";
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";

import type {
  AssetGetOKResponse,
  AssetUploadBody,
  AssetUploadOKResponse,
  InternalServerErrorResponse,
  NotFoundResponse,
  UnauthorisedResponse,
} from "./schemas";

/**
 * Upload and process a media file.
 */
export const assetUpload = (assetUploadBody: AssetUploadBody) => {
  return fetcher<AssetUploadOKResponse>({
    url: `/v1/assets`,
    method: "POST",
    headers: { "Content-Type": "application/octet-stream" },
    data: assetUploadBody,
  });
};

export const getAssetUploadMutationFetcher = () => {
  return (
    _: string,
    { arg }: { arg: Arguments },
  ): Promise<AssetUploadOKResponse> => {
    return assetUpload(arg as AssetUploadBody);
  };
};
export const getAssetUploadMutationKey = () => `/v1/assets` as const;

export type AssetUploadMutationResult = NonNullable<
  Awaited<ReturnType<typeof assetUpload>>
>;
export type AssetUploadMutationError =
  | UnauthorisedResponse
  | InternalServerErrorResponse;

export const useAssetUpload = <
  TError = UnauthorisedResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof assetUpload>>,
    TError,
    string,
    Arguments,
    Awaited<ReturnType<typeof assetUpload>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getAssetUploadMutationKey();
  const swrFn = getAssetUploadMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Download an asset by its ID.
 */
export const assetGet = (id: string) => {
  return fetcher<AssetGetOKResponse>({
    url: `/v1/assets/${id}`,
    method: "GET",
  });
};

export const getAssetGetKey = (id: string) => [`/v1/assets/${id}`] as const;

export type AssetGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof assetGet>>
>;
export type AssetGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useAssetGet = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  id: string,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof assetGet>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!id;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getAssetGetKey(id) : null));
  const swrFn = () => assetGet(id);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};

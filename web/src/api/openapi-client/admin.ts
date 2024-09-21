/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { Arguments } from "swr";
import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration } from "swr/mutation";

import { fetcher } from "../client";
import type {
  AccountGetOKResponse,
  AdminSettingsUpdateBody,
  AdminSettingsUpdateOKResponse,
  BadRequestResponse,
  InternalServerErrorResponse,
  NotFoundResponse,
  UnauthorisedResponse,
} from "../openapi-schema";

/**
 * Update non-env configuration settings for installation.
 */
export const adminSettingsUpdate = (
  adminSettingsUpdateBody: AdminSettingsUpdateBody,
) => {
  return fetcher<AdminSettingsUpdateOKResponse>({
    url: `/admin`,
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    data: adminSettingsUpdateBody,
  });
};

export const getAdminSettingsUpdateMutationFetcher = () => {
  return (
    _: string,
    { arg }: { arg: AdminSettingsUpdateBody },
  ): Promise<AdminSettingsUpdateOKResponse> => {
    return adminSettingsUpdate(arg);
  };
};
export const getAdminSettingsUpdateMutationKey = () => `/admin` as const;

export type AdminSettingsUpdateMutationResult = NonNullable<
  Awaited<ReturnType<typeof adminSettingsUpdate>>
>;
export type AdminSettingsUpdateMutationError =
  | BadRequestResponse
  | UnauthorisedResponse
  | InternalServerErrorResponse;

export const useAdminSettingsUpdate = <
  TError =
    | BadRequestResponse
    | UnauthorisedResponse
    | InternalServerErrorResponse,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof adminSettingsUpdate>>,
    TError,
    string,
    AdminSettingsUpdateBody,
    Awaited<ReturnType<typeof adminSettingsUpdate>>
  > & { swrKey?: string };
}) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey = swrOptions?.swrKey ?? getAdminSettingsUpdateMutationKey();
  const swrFn = getAdminSettingsUpdateMutationFetcher();

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Suspend an account - soft delete. This disables the ability for the
account owner to log in and use the platform. It keeps the account on
record for linkage to content so UI doesn't break. It does not change
anything else about the account such as the avatar, name, etc.

 */
export const adminAccountBanCreate = (accountHandle: string) => {
  return fetcher<AccountGetOKResponse>({
    url: `/admin/bans/${accountHandle}`,
    method: "POST",
  });
};

export const getAdminAccountBanCreateMutationFetcher = (
  accountHandle: string,
) => {
  return (_: string, __: { arg: Arguments }): Promise<AccountGetOKResponse> => {
    return adminAccountBanCreate(accountHandle);
  };
};
export const getAdminAccountBanCreateMutationKey = (accountHandle: string) =>
  `/admin/bans/${accountHandle}` as const;

export type AdminAccountBanCreateMutationResult = NonNullable<
  Awaited<ReturnType<typeof adminAccountBanCreate>>
>;
export type AdminAccountBanCreateMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useAdminAccountBanCreate = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  accountHandle: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof adminAccountBanCreate>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof adminAccountBanCreate>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getAdminAccountBanCreateMutationKey(accountHandle);
  const swrFn = getAdminAccountBanCreateMutationFetcher(accountHandle);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
/**
 * Given the account is suspended, remove the suspended state.
 */
export const adminAccountBanRemove = (accountHandle: string) => {
  return fetcher<AccountGetOKResponse>({
    url: `/admin/bans/${accountHandle}`,
    method: "DELETE",
  });
};

export const getAdminAccountBanRemoveMutationFetcher = (
  accountHandle: string,
) => {
  return (_: string, __: { arg: Arguments }): Promise<AccountGetOKResponse> => {
    return adminAccountBanRemove(accountHandle);
  };
};
export const getAdminAccountBanRemoveMutationKey = (accountHandle: string) =>
  `/admin/bans/${accountHandle}` as const;

export type AdminAccountBanRemoveMutationResult = NonNullable<
  Awaited<ReturnType<typeof adminAccountBanRemove>>
>;
export type AdminAccountBanRemoveMutationError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useAdminAccountBanRemove = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  accountHandle: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof adminAccountBanRemove>>,
      TError,
      string,
      Arguments,
      Awaited<ReturnType<typeof adminAccountBanRemove>>
    > & { swrKey?: string };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const swrKey =
    swrOptions?.swrKey ?? getAdminAccountBanRemoveMutationKey(accountHandle);
  const swrFn = getAdminAccountBanRemoveMutationFetcher(accountHandle);

  const query = useSWRMutation(swrKey, swrFn, swrOptions);

  return {
    swrKey,
    ...query,
  };
};
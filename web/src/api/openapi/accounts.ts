/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { SWRConfiguration, Key } from "swr";
import type {
  AccountsGetSuccessResponse,
  UnauthorisedResponse,
  NotFoundResponse,
  InternalServerErrorResponse,
  AccountsSetAvatarBody,
  AccountsGetAvatarResponse,
} from "./schemas";
import { fetcher } from "../client";

/**
 * @summary Get the information for the currently authenticated account.
 */
export const accountsGet = () => {
  return fetcher<AccountsGetSuccessResponse>({
    url: `/v1/accounts`,
    method: "get",
  });
};

export const getAccountsGetKey = () => [`/v1/accounts`];

export type AccountsGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof accountsGet>>
>;
export type AccountsGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useAccountsGet = <
  TError = UnauthorisedResponse | NotFoundResponse | InternalServerErrorResponse
>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof accountsGet>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getAccountsGetKey() : null));
  const swrFn = () => accountsGet();

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions
  );

  return {
    swrKey,
    ...query,
  };
};

/**
 * @summary Upload an avatar for the authenticated account.
 */
export const accountsSetAvatar = (
  accountsSetAvatarBody: AccountsSetAvatarBody
) => {
  return fetcher<void>({
    url: `/v1/accounts/self/avatar`,
    method: "post",
    headers: { "Content-Type": "application/octet-stream" },
    data: accountsSetAvatarBody,
  });
};

/**
 * @summary Get an avatar for the authenticated account.
 */
export const accountsGetAvatar = (accountId: string) => {
  return fetcher<AccountsGetAvatarResponse>({
    url: `/v1/accounts/${accountId}/avatar`,
    method: "get",
  });
};

export const getAccountsGetAvatarKey = (accountId: string) => [
  `/v1/accounts/${accountId}/avatar`,
];

export type AccountsGetAvatarQueryResult = NonNullable<
  Awaited<ReturnType<typeof accountsGetAvatar>>
>;
export type AccountsGetAvatarQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useAccountsGetAvatar = <
  TError = UnauthorisedResponse | NotFoundResponse | InternalServerErrorResponse
>(
  accountId: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof accountsGetAvatar>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  }
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!accountId;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getAccountsGetAvatarKey(accountId) : null));
  const swrFn = () => accountsGetAvatar(accountId);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions
  );

  return {
    swrKey,
    ...query,
  };
};

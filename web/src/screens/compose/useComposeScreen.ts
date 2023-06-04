import { useEffect, useState } from "react";

import { Thread } from "src/api/openapi/schemas";
import { threadGet } from "src/api/openapi/threads";
import { errorToast } from "src/components/ErrorBanner";

export type Props = { editing?: string };

export function useComposeScreen({ editing }: Props) {
  const [loadingDraft, setLoadingDraft] = useState(editing !== undefined);
  const [draft, setDraft] = useState<Thread | undefined>(undefined);

  useEffect(() => {
    const getDraft = async () => {
      if (editing === undefined) return;

      const thread = await threadGet(editing);

      setDraft(thread);

      return;
    };

    getDraft()
      .catch(errorToast)
      .finally(() => setLoadingDraft(false));
  }, [editing]);

  return {
    loadingDraft,
    draft,
  };
}

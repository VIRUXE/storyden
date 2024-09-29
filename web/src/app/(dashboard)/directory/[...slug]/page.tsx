import { notFound, redirect } from "next/navigation";

import { getServerSession } from "src/auth/server-session";
import { getTargetSlug } from "src/components/directory/datagraph/utils";
import { NodeCreateScreen } from "src/screens/directory/datagraph/NodeCreateScreen/NodeCreateScreen";
import { NodeViewerScreen } from "src/screens/directory/datagraph/NodeViewerScreen/NodeViewerScreen";
import {
  Params,
  ParamsSchema,
  Query,
} from "src/screens/directory/datagraph/directory-path";

import { nodeGet } from "@/api/openapi-server/nodes";

type Props = {
  params: Params;
  searchParams: Query;
};

export default async function Page(props: Props) {
  const { slug } = ParamsSchema.parse(props.params);
  const session = await getServerSession();

  const [targetSlug, isNew] = getTargetSlug(slug);

  if (targetSlug) {
    if (isNew) {
      if (!session) {
        redirect(`/login`); // TODO: ?return= back to this path.
      }

      return <NodeCreateScreen session={session} />;
    }

    const { data } = await nodeGet(targetSlug);

    if (data) {
      return <NodeViewerScreen slug={targetSlug} node={data} />;
    }
  }

  // Creating a new item or node from the root: "/directory/new"
  if (isNew) {
    if (!session) {
      redirect(`/login`); // TODO: ?return= back to this path.
    }

    return <NodeCreateScreen session={session} />;
  }

  notFound();
}

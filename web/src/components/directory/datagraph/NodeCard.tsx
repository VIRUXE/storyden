import { Node } from "src/api/openapi/schemas";
import { Timestamp } from "src/components/site/Timestamp";
import {
  DirectoryPath,
  joinDirectoryPath,
} from "src/screens/directory/datagraph/directory-path";
import { Card } from "src/theme/components/Card";

import { DirectoryBadge } from "../DirectoryBadge";

import { HStack } from "@/styled-system/jsx";
import { CardVariantProps } from "@/styled-system/recipes";

export type NodeCardContext = "directory" | "generic";

export type Props = {
  node: Node;
  directoryPath: DirectoryPath;
  context: NodeCardContext;
} & CardVariantProps;

export function NodeCard({ node, directoryPath, context, ...rest }: Props) {
  const slug = joinDirectoryPath(directoryPath, node.slug);
  const asset = node.assets?.[0];
  const url = `/directory/${slug}`;

  return (
    <Card
      id={node.id}
      title={node.name}
      text={node.description}
      url={url}
      image={asset?.url}
      {...rest}
    >
      {context === "generic" ? (
        <HStack color="fg.muted">
          <DirectoryBadge />

          <Timestamp
            created={node.createdAt}
            updated={node.updatedAt}
            href={url}
          />
        </HStack>
      ) : (
        <Timestamp
          created={node.createdAt}
          updated={node.updatedAt}
          large
          href={url}
        />
      )}
    </Card>
  );
}
import { Input } from "@chakra-ui/react";

import { Thread } from "src/api/openapi/schemas";

import { CategoryPill } from "../CategoryPill";

import { styled } from "@/styled-system/jsx";

import { useTitle } from "./useTitle";

export function Title(props: Thread) {
  const { editing, editingTitle, onTitleChange } = useTitle(props);

  return (
    <>
      <div>
        {editing ? (
          <Input value={editingTitle} onChange={onTitleChange} />
        ) : (
          <styled.h1 fontSize="var(--font-size-h1-variable)" fontWeight={600}>
            {props.title}
          </styled.h1>
        )}
      </div>
      <CategoryPill category={props.category} />
    </>
  );
}

import { useCategoriesQuery, useCategoryMutation } from "@/hooks/useCategory";

export const useSelectOptions = (type) => {
  switch (type) {
    case "category": {
      const { data = {}, ...rest } = useCategoriesQuery();
      return {
        data: data?.data || [],
        ...rest,
      };
    }
    default:
      throw new Error(`Unknown select type: ${type}`);
  }
};
export const useMutationOptions = (type) => {
  switch (type) {
    case "category":
      return useCategoryMutation();
    default:
      throw new Error(`Unknown select type: ${type}`);
  }
};

import { create } from "zustand";

export const useUserStore = create((set, get) => ({
  q: "",
  page: 1,
  limit: 10,
  sort: "created_at_desc",

  setQ: (val) => set({ q: val, page: 1 }),
  setPage: (val) => set({ page: val }),
  setLimit: (val) => set({ limit: val }),
  setSort: (field) => {
    const currentSort = get().sort;
    const newSort =
      currentSort === `${field}_asc` ? `${field}_desc` : `${field}_asc`;
    set({ sort: newSort, page: 1 });
  },

  reset: () =>
    set({
      q: "",
      page: 1,
      limit: 10,
      sort: "created_at_desc",
    }),
}));

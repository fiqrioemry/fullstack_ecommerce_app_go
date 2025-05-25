import { create } from "zustand";

export const useOrderStore = create((set, get) => ({
  q: "",
  page: 1,
  limit: 10,
  sort: "created_at_desc",
  status: "all",

  setQ: (val) => set({ q: val, page: 1 }),
  setPage: (val) => set({ page: val }),
  setLimit: (val) => set({ limit: val }),
  setStatus: (val) => set({ status: val, page: 1 }),

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
      status: "all",
    }),
}));

import { create } from "zustand";

export const useTransactionStore = create((set, get) => ({
  q: "",
  page: 1,
  limit: 10,
  sort: "paid_at_desc",
  status: "all",

  setPage: (val) => set({ page: val }),
  setLimit: (val) => set({ limit: val }),
  setQ: (val) => set({ q: val, page: 1 }),
  setStatus: (val) => set({ status: val }),
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
      sort: "paid_at_desc",
      status: "all",
    }),
}));

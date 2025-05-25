import { create } from "zustand";

export const useProductStore = create((set) => ({
  q: "",
  page: 1,
  limit: 10,
  sort: "",
  status: "all",
  minPrice: null,
  maxPrice: null,
  category: "",
  rating: null,
  range: "daily",

  setRange: (val) => set({ range: val }),
  setRating: (val) => set({ rating: val }),
  setCategory: (val) => set({ category: val }),
  setMinPrice: (val) => set({ minPrice: val }),
  setMaxPrice: (val) => set({ maxPrice: val }),
  setPage: (val) => set({ page: val }),
  setStatus: (val) => set({ status: val }),
  setLimit: (val) => set({ limit: val }),
  setQ: (val) => set({ q: val, page: 1 }),
  setSort: (field) =>
    set((state) => ({
      sort: state.sort === `${field}_asc` ? `${field}_desc` : `${field}_asc`,
      page: 1,
    })),

  reset: () =>
    set({
      q: "",
      page: 1,
      limit: 10,
      sort: "",
      status: "all",
      minPrice: null,
      maxPrice: null,
      category: "",
      rating: null,
      range: "daily",
    }),
}));

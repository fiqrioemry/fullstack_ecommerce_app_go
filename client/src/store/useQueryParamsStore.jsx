import { create } from "zustand";

export const useQueryParamsStore = create((set) => ({
  page: 1,
  limit: 10,
  status: "",
  search: "",
  categoryId: "",
  minPrice: null,
  maxPrice: null,
  sort: "",

  setPage: (page) =>
    set((state) => ({
      page: typeof page === "function" ? page(state.page) : page,
    })),

  setLimit: (limit) =>
    set((state) => ({
      limit: typeof limit === "function" ? limit(state.limit) : limit,
    })),

  setSort: (sort) =>
    set((state) => ({
      sort: typeof sort === "function" ? sort(state.sort) : sort,
      page: 1,
    })),

  setSearch: (search) =>
    set((state) => ({
      search: typeof search === "function" ? search(state.search) : search,
      page: 1,
    })),

  setStatus: (status) =>
    set((state) => ({
      status: typeof status === "function" ? status(state.status) : status,
      page: 1,
    })),

  setMinPrice: (minPrice) =>
    set((state) => ({
      minPrice:
        typeof minPrice === "function" ? minPrice(state.minPrice) : minPrice,
      page: 1,
    })),

  setMaxPrice: (maxPrice) =>
    set((state) => ({
      maxPrice:
        typeof maxPrice === "function" ? maxPrice(state.maxPrice) : maxPrice,
      page: 1,
    })),

  setCategoryId: (categoryId) =>
    set((state) => ({
      categoryId:
        typeof categoryId === "function"
          ? categoryId(state.categoryId)
          : categoryId,
      page: 1,
    })),

  reset: () =>
    set(() => ({
      search: "",
      page: 1,
      limit: 10,
      status: "",
      categoryId: "",
      minPrice: null,
      maxPrice: null,
      sort: "created_at desc",
    })),
}));

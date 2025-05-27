import { create } from "zustand";

export const useCheckoutStore = create((set) => ({
  note: "",
  courier: "sicepat",
  selectedShipping: null,
  shippingOptions: [],
  voucherCode: "",
  voucherInfo: null,

  setNote: (note) => set({ note }),
  setCourier: (courier) => set({ courier }),
  setVoucherCode: (code) => set({ voucherCode: code }),
  setVoucherInfo: (info) => set({ voucherInfo: info }),
  setSelectedShipping: (shipping) => set({ selectedShipping: shipping }),
  setShippingOptions: (options) => set({ shippingOptions: options }),

  resetCheckout: () =>
    set({
      note: "",
      courier: "sicepat",
      selectedShipping: null,
      shippingOptions: [],
      voucherCode: "",
      voucherInfo: null,
    }),
}));

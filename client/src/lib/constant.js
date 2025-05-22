// AUTHENTICATION
export const registerState = {
  email: "",
  password: "",
  fullname: "",
  otp: "",
};

export const sendOTPState = {
  email: "",
};

export const verifyOTPState = {
  email: "",
  otp: "",
};
export const getLoginState = (rememberMe = false) => ({
  email: "",
  password: "",
  rememberMe,
});

// PROFILE MANAGEMENT
export const profileState = {
  fullname: "",
  birthday: "",
  gender: "",
  phone: "",
};

export const addressState = {
  name: "",
  address: "",
  provinceId: 0,
  cityId: 0,
  districtId: 0,
  subdistrictId: 0,
  postalCodeId: 0,
  phone: "",
  isMain: false,
};

export const genderOptions = [
  { value: "male", label: "Male" },
  { value: "female", label: "Female" },
];

// PRODUCT, CATEGORY, BANNER
export const productState = {
  name: "",
  description: "",
  price: 0,
  stock: 0,
  discount: 0,
  weight: 1000,
  height: 0,
  width: 0,
  length: 0,
  isFeatured: false,
  isActive: true,
  images: undefined,
  categoryId: "",
};

export const bannerState = {
  position: "",
  image: undefined,
};

export const categoryState = {
  name: "",
  image: undefined,
};

// PAYMENT AND SHIPMENT
export const midtransNotificationState = {
  transaction_status: "",
  order_id: "",
  payment_type: "",
  fraud_status: "",
};

export const paymentState = {
  packageId: "",
};

export const shipmentState = {
  trackingCode: "",
  note: "",
};

// NOTIFICATION AND VOUCHER
export const typeCode = [
  { label: "System Message", value: "system_message" },
  { label: "Class Reminder", value: "class_reminder" },
  { label: "Promo Offer", value: "promo_offer" },
];

export const createVoucherState = {
  code: "",
  description: "",
  discountType: "fixed",
  discount: 0,
  maxDiscount: null,
  quota: 1,
  expiredAt: "",
};

export const notificationState = {
  title: "",
  message: "",
  typeCode: "",
};

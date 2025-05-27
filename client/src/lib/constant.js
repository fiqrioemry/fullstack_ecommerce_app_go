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

export const reviewState = {
  rating: "",
  image: undefined,
  comment: "",
};

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

// FILTER OPTIONS
export const orderStatusOptions = [
  { value: "all", label: "All" },
  { value: "waiting_payment", label: "Waiting payment" },
  { value: "success", label: "Success" },
  { value: "pending", label: "Pending" },
  { value: "process", label: "Process" },
  { value: "canceled", label: "Canceled" },
];

export const paymentStatusOptions = [
  { value: "all", label: "All" },
  { value: "success", label: "Success" },
  { value: "pending", label: "Pending" },
  { value: "failed", label: "failed" },
];

export const productSortOptions = [
  { value: "name_asc", label: "Name A-Z" },
  { value: "name_desc", label: "Name Z-A" },
  { value: "price_asc", label: "Price Low to High" },
  { value: "price_desc", label: "Price High to Low" },
  { value: "created_at_asc", label: "Latest" },
  { value: "created_at_desc", label: "Oldest" },
];

export const paymentSortOptions = [
  { value: "paid_at_desc", label: "Newest Payment" },
  { value: "paid_at_asc", label: "Oldest Payment" },
  { value: "total_desc", label: "Highest Total" },
  { value: "total_asc", label: "Lowest Total" },
];

export const orderSortOptions = [
  { value: "created_at_desc", label: "Newest Order" },
  { value: "created_at_asc", label: "Oldest Order" },
  { value: "product_name_asc", label: "Product Name A-Z" },
  { value: "product_name_desc", label: "Product Name Z-A" },
];

export const productStatusOptions = [
  { value: "all", label: "All" },
  { value: "active", label: "Active" },
  { value: "inactive", label: "Inactive" },
  { value: "featured", label: "Featured" },
  { value: "unfeatured", label: "Not Featured" },
];

export const revenueRangeOptions = [
  { value: "daily", label: "Daily" },
  { value: "monthly", label: "monthly" },
  { value: "yearly", label: "yearly" },
];

export const ratingOptions = [
  { value: "4", label: "4 and above" },
  { value: "3", label: "3 and above" },
  { value: "2", label: "2 and above" },
  { value: "1", label: "1 and above" },
];

export const typeCode = [
  { label: "Promo Offer", value: "promo_offer" },
  { label: "System Message", value: "system_message" },
  { label: "Class Reminder", value: "class_reminder" },
];

export const genderOptions = [
  { value: "male", label: "Male" },
  { value: "female", label: "Female" },
];

export const courierOptions = [
  { value: "jne", label: "JNE" },
  { value: "sicepat", label: "SiCepat" },
];

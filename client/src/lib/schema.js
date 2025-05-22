import { z } from "zod";

const imageItemSchema = z
  .any()
  .optional()
  .refine((file) => !file || file instanceof File || typeof file === "string", {
    message: "Input must be a file or a valid URL",
  })
  .refine(
    (file) =>
      !file || typeof file === "string" || file.type?.startsWith("image/"),
    { message: "File must be an image" }
  )
  .refine(
    (file) => !file || typeof file === "string" || file.size <= 2 * 1024 * 1024,
    { message: "Image size must be <= 2MB" }
  );

export const sendOTPSchema = z.object({
  email: z.string().email("Invalid email address"),
});

export const verifyOTPSchema = z.object({
  email: z.string().email("Invalid email address"),
  otp: z.string().min(6, "OTP code must be at least 6 characters"),
});

export const registerSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  fullname: z.string().min(1, "Full name is required"),
});

export const loginSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  rememberMe: z.boolean().optional(),
});

// PROFILE MANAGEMENT
export const profileSchema = z.object({
  fullname: z.string().min(6, "Fullname must be at least 6 characters"),
  birthday: z.string().refine((val) => !isNaN(Date.parse(val)), {
    message: "Tanggal tidak valid",
  }),
  gender: z.string().optional(),
  phone: z.string().optional(),
  bio: z.string().optional(),
});

export const avatarSchema = z.object({
  avatar: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
});

export const addressSchema = z.object({
  name: z.string().min(1, "Name is required"),
  address: z.string().min(1, "Address is required"),
  provinceId: z.number().min(1, "Province is required"),
  cityId: z.number().min(1, "City is required"),
  districtId: z.number().min(1, "District is required"),
  subdistrictId: z.number().min(1, "Subdistrict is required"),
  postalCodeId: z.number().min(1, "Postal Code is required"),
  phone: z.string().min(8, "Phone is required"),
  isMain: z.boolean().optional(),
});

// PRODUCT, CATEGORY, BANNER
export const productSchema = z.object({
  name: z.string().min(5, { message: "Name must be at least 5 characters" }),
  categoryId: z.string().min(1, "Category is required"),
  description: z
    .string()
    .min(20, { message: "Description must be at least 20 characters" }),
  price: z
    .number({ invalid_type_error: "Price is required" })
    .positive({ message: "Price must be greater than 0" }),
  stock: z
    .number({ invalid_type_error: "Stock is required" })
    .min(0, { message: "Stock must be at least 0" }),
  discount: z
    .number()
    .min(0, "Discount cannot be negative")
    .max(100, "Discount cannot exceed 100")
    .optional(),
  isActive: z.boolean().optional(),
  isFeatured: z.boolean().optional(),
  weight: z
    .number({ invalid_type_error: "Weight is required" })
    .min(0, { message: "Minimum required for shipment 1000 gr" }),
  length: z.number({ invalid_type_error: "Length is required" }),
  width: z.number({ invalid_type_error: "Width is required" }),
  height: z.number({ invalid_type_error: "Height is required" }),
  images: z.array(imageItemSchema).min(1, "Image is required"),
});
export const bannerSchema = z.object({
  position: z.string().min(1, { message: "Position for banner is required" }),
  image: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
});

export const categorySchema = z.object({
  name: z
    .string()
    .min(5, { message: "Category must be at least 5 characters" }),
  image: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
});

// PAYMENT AND SHIPMENT
export const createPaymentSchema = z.object({
  packageId: z.string().min(1, "Package is required"),
});
export const midtransNotificationSchema = z.object({
  transaction_status: z.string(),
  order_id: z.string(),
  payment_type: z.string(),
  fraud_status: z.string(),
});

export const shipmentSchema = z.object({
  trackingCode: z.string().min(1, "Tracking number is required"),
  note: z.string().optional(),
});

// NOTIFICATIONS AND VOUCHERS

export const createReviewSchema = z.object({
  classId: z.string().min(1, "Class is required"),
  rating: z
    .number({
      required_error: "Rating is required",
      invalid_type_error: "Rating must be a number",
    })
    .min(1, "Minimum rating is 1")
    .max(5, "Maximum rating is 5"),
  comment: z.string().min(8, "Comment must be at least 8 characters"),
});

export const createVoucherSchema = z
  .object({
    code: z.string().min(1, "Code is required"),
    description: z.string().min(1, "Description is required"),
    discountType: z.enum(["fixed", "percentage"], {
      required_error: "Discount type is required",
      invalid_type_error: "Please select a valid discount type",
    }),
    discount: z
      .number({ invalid_type_error: "Discount must be a number" })
      .gt(0, "Discount must be greater than 0"),
    maxDiscount: z
      .number({ invalid_type_error: "Max discount must be a number" })
      .optional()
      .nullable(),
    quota: z
      .number({ invalid_type_error: "Quota must be a number" })
      .gt(0, "Quota must be greater than 0"),
    expiredAt: z.string().refine((val) => !isNaN(Date.parse(val)), {
      message: "Expired date must be valid (YYYY-MM-DD)",
    }),
  })
  .superRefine((val, ctx) => {
    if (val.discountType === "percentage" && val.discount > 100) {
      ctx.addIssue({
        code: "custom",
        path: ["discount"],
        message: "Percentage discount cannot exceed 100",
      });
    }
  });

export const notificationSchema = z.object({
  title: z.string().min(3, "Title is required"),
  message: z
    .string()
    .min(5, "Message is required")
    .max(200, "Maximum 200 characters allowed"),
  typeCode: z.enum(["system_message", "class_reminder", "promo_offer"]),
});

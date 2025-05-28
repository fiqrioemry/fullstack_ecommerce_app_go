// src/lib/utils.js
import {
  format,
  parse,
  getDay,
  formatDuration,
  intervalToDuration,
} from "date-fns";
import { clsx } from "clsx";
import id from "date-fns/locale/id";
import { twMerge } from "tailwind-merge";
import { dateFnsLocalizer } from "react-big-calendar";

export function cn(...inputs) {
  return twMerge(clsx(inputs));
}

export const formatRupiah = (number) => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    maximumFractionDigits: 0,
  }).format(number);
};

export const formatDateTime = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleString("en-GB", {
    year: "numeric",
    month: "short",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  });
};

export const formatDate = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleDateString("en-GB", {
    year: "numeric",
    month: "short",
    day: "2-digit",
  });
};

export const buildFormData = (data) => {
  const formData = new FormData();

  Object.entries(data).forEach(([key, value]) => {
    if (value === undefined || value === null) return;

    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof File) {
        value.forEach((file) => {
          formData.append(key, file);
        });
      } else {
        value.forEach((item) => {
          formData.append(`${key}`, item);
        });
      }
    } else if (value instanceof File) {
      formData.append(key, value);
    } else {
      formData.append(key, value);
    }
  });

  return formData;
};

export const truncateText = (text, maxLength) => {
  return text.length > maxLength ? `${text.slice(0, maxLength)}...` : text;
};

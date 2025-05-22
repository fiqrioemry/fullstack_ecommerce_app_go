import {
  CreditCard,
  MailOpen,
  MapPin,
  Settings2Icon,
  UserRoundPen,
} from "lucide-react";
import { cn } from "@/lib/utils";
import { Fragment } from "react";
import Header from "../public/Header";
import Footer from "../public/Footer";
import { MobileSidebar } from "./MobileSidebar";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, Outlet, useLocation } from "react-router-dom";

const customerMenu = [
  {
    title: "profile",
    path: "/user/profile",
    icon: UserRoundPen,
  },
  {
    title: "address",
    path: "/user/addresses",
    icon: MapPin,
  },
  {
    title: "transactions",
    path: "/user/transactions",
    icon: CreditCard,
  },
  {
    title: "noitifications",
    path: "/user/notifications",
    icon: MailOpen,
  },
  {
    title: "settings",
    path: "/user/settings",
    icon: Settings2Icon,
  },
];

const CustomerLayout = () => {
  const { user } = useAuthStore();
  const location = useLocation();
  const currentPath = location.pathname;

  return (
    <Fragment>
      <Header />
      <MobileSidebar />
      <main className="py-20 container max-w-7xl mx-auto relative">
        {/* Floating Sidebar */}
        <div className="absolute left-0 top-0 bottom-0 hidden md:block">
          <div className="bg-background shadow-lg rounded-xl w-64 p-6 mt-10 ml-4 sticky top-24 flex flex-col items-center">
            {/* Avatar Bulat */}
            <img
              src={user.avatar}
              alt={user.fullname}
              className="w-20 h-20 rounded-full border-4 border-white shadow-md mb-6"
            />

            {/* Menu */}
            <div className="w-full flex flex-col gap-2">
              {customerMenu.map((menu) => {
                const activePath = currentPath === menu.path;
                return (
                  <Link
                    to={menu.path}
                    key={menu.title}
                    className={cn(
                      "flex items-center gap-3 p-2 rounded-lg transition hover:bg-accent",
                      activePath
                        ? "bg-accent font-semibold text-primary"
                        : "text-muted-foreground"
                    )}
                  >
                    <menu.icon className="w-5 h-5" />
                    <span className="capitalize">{menu.title}</span>
                  </Link>
                );
              })}
            </div>
          </div>
        </div>

        {/* Content */}
        <div className="md:ml-72 mt-2">
          <Outlet />
        </div>
      </main>

      <Footer />
    </Fragment>
  );
};

export default CustomerLayout;

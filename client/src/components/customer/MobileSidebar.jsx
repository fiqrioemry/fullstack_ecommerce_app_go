import { cn } from "@/lib/utils";
import { Menu } from "lucide-react";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, useLocation } from "react-router-dom";
import { Sheet, SheetTrigger, SheetContent } from "@/components/ui/sheet";

const customerMenu = [
  { title: "profile", path: "/user/profile", icon: "✏️" },
  { title: "address", path: "/user/address", icon: "📍" },
  { title: "transactions", path: "/user/transactions", icon: "💳" },
  { title: "orders", path: "/user/orders", icon: "🚚" },
  { title: "noitifications", path: "/user/notifications", icon: "📨" },
  { title: "settings", path: "/user/settings", icon: "⚙️" },
];

const MobileSidebar = () => {
  const location = useLocation();
  const currentPath = location.pathname;
  const { user } = useAuthStore();

  return (
    <Sheet>
      <SheetTrigger className="md:hidden fixed top-14 left-4 z-50">
        <Menu className="w-6 h-6" />
      </SheetTrigger>
      <SheetContent side="left" className="w-[240px] p-5">
        <div className="flex flex-col items-center">
          <img
            src={
              user?.photoURL ||
              `https://api.dicebear.com/6.x/initials/svg?seed=${
                user?.displayName || "User"
              }`
            }
            alt="avatar"
            className="w-16 h-16 rounded-full mb-4 border"
          />
          <div className="w-full flex flex-col gap-2">
            {customerMenu.map((menu) => (
              <Link
                key={menu.path}
                to={menu.path}
                className={cn(
                  "flex items-center gap-3 p-2 rounded-lg transition hover:bg-accent",
                  currentPath === menu.path
                    ? "bg-accent font-semibold text-primary"
                    : "text-muted-foreground"
                )}
              >
                <span>{menu.icon}</span>
                <span className="capitalize">{menu.title}</span>
              </Link>
            ))}
          </div>
        </div>
      </SheetContent>
    </Sheet>
  );
};

export { MobileSidebar };

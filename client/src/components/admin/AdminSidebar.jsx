import {
  Sidebar,
  SidebarMenu,
  SidebarHeader,
  SidebarFooter,
  SidebarContent,
} from "@/components/ui/sidebar";
import {
  DropdownMenu,
  DropdownMenuItem,
  DropdownMenuTrigger,
  DropdownMenuContent,
} from "@/components/ui/dropdown-menu";
import { cn } from "@/lib/utils";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, useLocation } from "react-router-dom";
import {
  Users,
  BarChart2,
  LogOut,
  BoxIcon,
  CreditCard,
  Truck,
} from "lucide-react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

const NavItem = ({ to, icon: Icon, title, active }) => (
  <Link
    to={to}
    className={cn(
      "flex items-center gap-3 px-4 py-2 text-sm rounded-md transition",
      active
        ? "bg-primary text-primary-foreground font-semibold"
        : "text-muted-foreground hover:bg-muted"
    )}
  >
    <Icon className="w-4 h-4" />
    {title}
  </Link>
);
const directMenus = [
  {
    to: "/admin/dashboard",
    icon: Users,
    title: "Users",
  },
  {
    to: "/admin/products",
    icon: BoxIcon,
    title: "Products",
  },
  {
    to: "/admin/orders",
    icon: Truck,
    title: "Orders",
  },
  {
    to: "/admin/transactions",
    icon: CreditCard,
    title: "Transactions",
  },
];

const AdminSidebar = () => {
  const location = useLocation();
  const currentPath = location.pathname;
  const { user, logout } = useAuthStore();

  return (
    <Sidebar>
      <SidebarContent className="px-4 space-y-4 text-sm text-gray-700">
        <SidebarHeader className="mb-4 py-2">
          <img src="/logo.png" />
        </SidebarHeader>

        <SidebarMenu className="space-y-1">
          <NavItem
            to="/admin"
            title="Dashboard"
            icon={BarChart2}
            active={currentPath === "/admin"}
          />

          {directMenus.map((item) => (
            <NavItem
              key={item.to}
              to={item.to}
              icon={item.icon}
              title={item.title}
              active={currentPath === item.to}
            />
          ))}
        </SidebarMenu>
      </SidebarContent>
      <SidebarFooter className="p-4 text-xs text-muted-foreground">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <div className="flex items-center gap-3 cursor-pointer hover:bg-muted px-3 py-2 rounded-md transition">
              <Avatar className="w-9 h-9">
                <AvatarImage src={user?.avatar} alt={user?.fullname} />
                <AvatarFallback>{user?.fullname?.[0] || "A"}</AvatarFallback>
              </Avatar>
              <div className="flex flex-col text-left overflow-hidden">
                <span className="text-sm font-medium text-foreground truncate">
                  {user?.fullname || "Admin"}
                </span>
                <span className="text-xs text-muted-foreground truncate">
                  {user?.email || "admin@gmail.com"}
                </span>
              </div>
            </div>
          </DropdownMenuTrigger>

          <DropdownMenuContent side="top" align="start" className="w-60">
            <DropdownMenuItem onClick={logout}>
              <LogOut className="w-4 h-4 mr-2" />
              Logout
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarFooter>
    </Sidebar>
  );
};

export default AdminSidebar;

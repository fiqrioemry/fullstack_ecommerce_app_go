import {
  Sidebar,
  SidebarMenu,
  SidebarHeader,
  SidebarFooter,
  SidebarContent,
} from "@/components/ui/Sidebar";
import {
  Accordion,
  AccordionItem,
  AccordionTrigger,
  AccordionContent,
} from "@/components/ui/accordion";

import {
  DropdownMenu,
  DropdownMenuItem,
  DropdownMenuTrigger,
  DropdownMenuContent,
} from "@/components/ui/DropdownMenu";
import { cn } from "@/lib/utils";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, useLocation } from "react-router-dom";
import {
  Users,
  BarChart2,
  LogOut,
  CreditCard,
  Truck,
  Package,
  List,
  Plus,
  Image,
  SquareM,
  UserIcon,
  Ticket,
  Mail,
} from "lucide-react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/Avatar";

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
    icon: BarChart2,
    title: "Dashboard",
  },
  {
    to: "/admin/users",
    icon: Users,
    title: "Users",
  },
  {
    to: "/admin/orders",
    icon: Truck,
    title: "Orders",
  },
  {
    to: "/admin/banners",
    icon: Image,
    title: "Banners",
  },
  {
    to: "/admin/categories",
    icon: SquareM,
    title: "Categories",
  },
  {
    to: "/admin/vouchers",
    icon: Ticket,
    title: "Vouchers",
  },
  {
    to: "/admin/messages",
    icon: Mail,
    title: "Messages",
  },
  {
    to: "/admin/transactions",
    icon: CreditCard,
    title: "Transactions",
  },
];

const accordionMenus = [
  {
    value: "products",
    icon: Package,
    title: "Products",
    children: [
      { to: "/admin/products", title: "Product List", icon: List },
      { to: "/admin/products/add", title: "Add Product", icon: Plus },
    ],
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
          {directMenus.map((item) => (
            <NavItem
              key={item.to}
              to={item.to}
              icon={item.icon}
              title={item.title}
              active={currentPath === item.to}
            />
          ))}
          <Accordion type="multiple" className="space-y-1">
            {accordionMenus.map((menu) => (
              <AccordionItem key={menu.value} value={menu.value}>
                <AccordionTrigger
                  className={cn(
                    "w-full px-4 py-2 text-sm rounded-md transition flex items-center gap-2",
                    "text-muted-foreground hover:bg-muted [&[data-state=open]]:bg-muted"
                  )}
                >
                  <menu.icon className="w-4 h-4" />
                  {menu.title}
                </AccordionTrigger>

                <AccordionContent className="pl-6 space-y-1 mt-1">
                  {menu.children.map((child) => (
                    <NavItem
                      key={child.to}
                      to={child.to}
                      title={child.title}
                      icon={child.icon}
                      active={currentPath === child.to}
                    />
                  ))}
                </AccordionContent>
              </AccordionItem>
            ))}
          </Accordion>
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
            <DropdownMenuItem asChild>
              <div className="flex items-center gap-3 cursor-pointer hover:bg-muted px-3 py-2 rounded-md transition">
                <UserIcon />
                <div className="flex flex-col text-left overflow-hidden">
                  <span className="text-sm font-medium text-foreground truncate">
                    {user?.fullname || "Admin"}
                  </span>
                </div>
              </div>
            </DropdownMenuItem>
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

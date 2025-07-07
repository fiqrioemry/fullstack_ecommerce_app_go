// public pages
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import NotFound from "./pages/NotFound";
import ProductDetail from "./pages/ProductDetail";
import ProductResults from "./pages/ProductResults";

// admin pages
import Messages from "./pages/admin/Messages";
import Dashboard from "./pages/admin/Dashboard";
import OrdersList from "./pages/admin/OrdersList";
import AddProduct from "./pages/admin/AddProduct";
import BannersList from "./pages/admin/BannersList";
import VouchersList from "./pages/admin/VoucherList";
import ProductsList from "./pages/admin/ProductsList";
import CustomersList from "./pages/admin/CustomersList";
import CategoriesList from "./pages/admin/CategoriesList";
import TransactionsList from "./pages/admin/TransactionsList";
import { OrderDetail } from "./components/admin/orders/OrderDetail";
import { CustomerDetail } from "./components/admin/users/CustomerDetail";

// customer pages
import Checkout from "./pages/Checkout";
import CartPage from "./pages/CartPage";
import InvoicePage from "./pages/InvoicePage";
import UserProfile from "./pages/customer/UserProfile";
import UserSettings from "./pages/customer/UserSettings";
import UserAddresses from "./pages/customer/UserAddresses";
import UserTransactions from "./pages/customer/UserTransactions";
import UserNotifications from "./pages/customer/UserNotifications";
import { TransactionDetail } from "./components/customer/transactions/TransactionDetail";

// route config & support
import "sonner/dist/sonner.css";
import { toast, Toaster } from "sonner";
import { useEffect } from "react";
import { Loading } from "@/components/ui/Loading";
import { useAuthStore } from "./store/useAuthStore";
import { ScrollToTop } from "./hooks/useScrollToTop";
import { Navigate, Route, Routes, useLocation } from "react-router-dom";
import { AdminRoute, AuthRoute, NonAuthRoute, PublicRoute } from "./middleware";

// pages layout
import AdminLayout from "./components/admin/AdminLayout";
import PublicLayout from "./components/public/PublicLayout";
import CustomerLayout from "./components/customer/CustomerLayout";

function App() {
  const location = useLocation();
  const state = location.state;
  const backgroundLocation = state?.backgroundLocation;
  const { checkingAuth, authMe } = useAuthStore();

  useEffect(() => {
    authMe();
  }, []);

  if (checkingAuth) return <Loading />;

  return (
    <>
      <Toaster position="top-center" />
      <ScrollToTop />
      <Routes location={backgroundLocation || location}>
        <Route
          path="/invoice/:orderId"
          element={
            <AuthRoute>
              <InvoicePage />
            </AuthRoute>
          }
        />
        <Route
          path="/signin"
          element={
            <NonAuthRoute>
              <SignIn />
            </NonAuthRoute>
          }
        />
        <Route
          path="/signup"
          element={
            <NonAuthRoute>
              <SignUp />
            </NonAuthRoute>
          }
        />
        {/* Public */}
        <Route
          path="/"
          element={
            <PublicRoute>
              <PublicLayout />
            </PublicRoute>
          }
        >
          <Route index element={<Home />} />
          <Route path="products" element={<ProductResults />} />
          <Route path="products/:slug" element={<ProductDetail />} />
          <Route
            path="cart"
            element={
              <AuthRoute>
                <CartPage />
              </AuthRoute>
            }
          />
          <Route
            path="cart/checkout"
            element={
              <AuthRoute>
                <Checkout />
              </AuthRoute>
            }
          />
        </Route>

        {/* customer */}
        <Route
          path="/user"
          element={
            <AuthRoute>
              <CustomerLayout />
            </AuthRoute>
          }
        >
          <Route path="profile" element={<UserProfile />} />
          <Route path="settings" element={<UserSettings />} />
          <Route path="addresses" element={<UserAddresses />} />
          <Route path="transactions" element={<UserTransactions />} />
          <Route path="notifications" element={<UserNotifications />} />

          <Route index element={<Navigate to="profile" replace />} />
        </Route>

        {/* admin */}
        <Route
          path="/admin"
          element={
            <AdminRoute>
              <AdminLayout />
            </AdminRoute>
          }
        >
          <Route path="users" element={<CustomersList />} />
          <Route path="orders" element={<OrdersList />} />
          <Route path="products" element={<ProductsList />} />
          <Route path="products/add" element={<AddProduct />} />
          <Route path="banners" element={<BannersList />} />
          <Route path="categories" element={<CategoriesList />} />
          <Route path="vouchers" element={<VouchersList />} />
          <Route path="messages" element={<Messages />} />
          <Route path="transactions" element={<TransactionsList />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route index element={<Navigate to="dashboard" replace />} />
        </Route>

        <Route path="*" element={<NotFound />} />
      </Routes>
      {/* background dialog */}
      {backgroundLocation && (
        <Routes>
          <Route path="/admin/users/:id" element={<CustomerDetail />} />
          <Route path="/admin/orders/:id" element={<OrderDetail />} />
          <Route
            path="/user/transactions/:id"
            element={<TransactionDetail />}
          />
        </Routes>
      )}
    </>
  );
}

export default App;

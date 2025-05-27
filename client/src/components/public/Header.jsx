import { useState } from "react";
import { LogIn } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { WebLogo } from "@/components/ui/WebLogo";
import { useDebounce } from "@/hooks/useDebounce";
import { useAuthStore } from "@/store/useAuthStore";
import { useSearchProductsQuery } from "@/hooks/useProduct";
import { CartDropdown } from "@/components/header/CartDropdown";
import { UserDropdown } from "@/components/header/UserDropdown";
import { SearchProduct } from "@/components/header/SearchProduct";
import { SearchDropdown } from "@/components/header/SearchDropdown";

const Header = () => {
  const [q, setQ] = useState("");
  const { user } = useAuthStore();
  const navigate = useNavigate();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading } = useSearchProductsQuery({
    q: debouncedQ,
    limit: 5,
  });

  const handleResultClick = (slug) => {
    navigate(`/products/${slug}`);
    setQ("");
  };

  const handleLoginClick = () => navigate("/signin");

  return (
    <header className="fixed w-full bg-background p-2 border-b shadow-sm z-50">
      <div className="section flex items-center justify-between gap-4">
        {/* Logo */}

        <WebLogo />
        {/* Search */}
        <div className="relative w-full max-w-md">
          <SearchProduct
            value={q}
            onChange={setQ}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                navigate(`/products?q=${q}`);
                setQ("");
              }
            }}
            onClear={() => setQ("")}
          />
          <SearchDropdown
            search={q}
            isLoading={isLoading}
            results={data?.data}
            onClick={handleResultClick}
          />
        </div>

        {/* Right section */}
        <div className="flex items-center gap-4">
          {/* Shopping cart dropdown*/}
          {user && <CartDropdown />}

          {user ? (
            <UserDropdown />
          ) : (
            <Button onClick={handleLoginClick}>
              <LogIn className="w-4 h-4" />
              Login
            </Button>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;

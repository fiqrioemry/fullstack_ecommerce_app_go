import { LogIn } from "lucide-react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { useSearchProductsQuery } from "@/hooks/useProduct";
import { SearchInput } from "@/components/header/SearchInput";
import { CartDropdown } from "@/components/header/CartDropdown";
import { UserDropdown } from "@/components/header/UserDropdown";
import { SearchDropdown } from "@/components/header/SearchDropdown";

const Header = () => {
  const { user } = useAuthStore();
  const navigate = useNavigate();
  const [search, setSearch] = useState("");
  const [debouncedSearch, setDebouncedSearch] = useState("");

  useEffect(() => {
    const handler = setTimeout(() => setDebouncedSearch(search), 500);
    return () => clearTimeout(handler);
  }, [search]);

  const { data, isFetching } = useSearchProductsQuery(
    debouncedSearch ? { q: debouncedSearch, limit: 5 } : null
  );

  const handleResultClick = (slug) => {
    navigate(`/products/${slug}`);
    setSearch("");
  };

  const handleLoginClick = () => navigate("/signin");

  return (
    <header className="fixed w-full bg-background p-2 border-b shadow-sm z-50">
      <div className="section flex items-center justify-between gap-4">
        {/* Logo */}

        <WebLogo />
        {/* Search */}
        <div className="relative w-full max-w-md">
          <SearchInput
            value={search}
            onChange={setSearch}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                navigate(`/products?q=${search}`);
                setSearch("");
              }
            }}
            onClear={() => setSearch("")}
          />
          <SearchDropdown
            search={search}
            isLoading={isFetching}
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

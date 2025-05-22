import {
  Select,
  SelectGroup,
  SelectLabel,
  SelectItem,
  SelectValue,
  SelectContent,
  SelectTrigger,
} from "@/components/ui/select";
import { useMemo, useState } from "react";
import { Grid2X2, List, X } from "lucide-react";
import { Loading } from "@/components/ui/Loading";
import { useSearchParams } from "react-router-dom";
import { Pagination } from "@/components/ui/pagination";
import { useCategoriesQuery } from "@/hooks/useCategory";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useSearchProductsQuery } from "@/hooks/useProduct";
import { ProductCard } from "@/components/product-results/ProductCard";
import { ProductList } from "@/components/product-results/ProductList";
import { NoProductResult } from "@/components/product-results/NoProductResult";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

const ProductResults = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [showPriceWarning, setShowPriceWarning] = useState(false);

  const queryParams = useMemo(() => {
    const params = Object.fromEntries(searchParams.entries());
    return {
      q: params.q || "",
      category: params.category || "",
      subcategory: params.subcategory || "",
      minPrice: parseFloat(params.minPrice) || undefined,
      maxPrice: parseFloat(params.maxPrice) || undefined,
      rating: parseFloat(params.rating) || undefined,
      sort: params.sort || "created_at desc",
      page: parseInt(params.page || "1"),
    };
  }, [searchParams]);

  const {
    data: searchData,
    isLoading,
    isError,
    refetch,
  } = useSearchProductsQuery({
    search: queryParams.q,
    page: queryParams.page,
    limit: 12,
    sort: queryParams.sort,
    category: queryParams.category,
    minPrice: queryParams.minPrice,
    maxPrice: queryParams.maxPrice,
    rating: queryParams.rating,
  });

  const { data: categoriesData, isLoading: isCategoriesLoading } =
    useCategoriesQuery();

  const categories = categoriesData?.data || [];
  const [minPriceInput, setMinPriceInput] = useState(
    queryParams.minPrice || ""
  );
  const [maxPriceInput, setMaxPriceInput] = useState(
    queryParams.maxPrice || ""
  );

  const handleFilterChange = (key, value) => {
    const params = new URLSearchParams(searchParams);
    if (value) params.set(key, value);
    else params.delete(key);
    if (params.get("page")) params.delete("page");
    setSearchParams(params);
  };

  const handlePageChange = (newPage) => {
    const params = new URLSearchParams(searchParams);
    if (newPage > 1) {
      params.set("page", newPage);
    } else {
      params.delete("page");
    }
    setSearchParams(params);
  };

  const handlePriceInputChange = (key, value) => {
    if (!/^[0-9]*$/.test(value)) return;
    key === "minPrice" ? setMinPriceInput(value) : setMaxPriceInput(value);
  };

  const applyPriceFilter = () => {
    const minVal = parseFloat(minPriceInput);
    const maxVal = parseFloat(maxPriceInput);

    setShowPriceWarning(minVal && maxVal && minVal > maxVal);

    const params = new URLSearchParams(searchParams);
    if (minPriceInput) params.set("minPrice", minPriceInput);
    else params.delete("minPrice");
    if (maxPriceInput) params.set("maxPrice", maxPriceInput);
    else params.delete("maxPrice");
    if (params.get("page")) params.delete("page");
    setSearchParams(params);
  };

  const activeFilters = Array.from(searchParams.entries()).filter(([key]) =>
    ["q", "category", "subcategory", "minPrice", "maxPrice", "rating"].includes(
      key
    )
  );

  const removeFilter = (key) => {
    const params = new URLSearchParams(searchParams);
    params.delete(key);
    if (params.get("page")) params.delete("page");
    setSearchParams(params);
  };

  const clearAllFilters = () => {
    setSearchParams({});
  };

  if (isLoading || isCategoriesLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const data = searchData?.data || [];
  const pagination = searchData?.pagination || {};

  return (
    <section className="section py-16 md:py-20 space-y-8">
      {showPriceWarning && (
        <div className="bg-red-100 text-red-600 text-sm p-3 rounded border border-red-300">
          Maximum price must be greater than minimum price.
        </div>
      )}

      <div className="grid grid-cols-4 gap-4">
        {/* Sidebar */}
        <div className="col-span-4 md:col-span-1 space-y-4">
          {/* Category */}
          <div className="space-y-2">
            <h3 className="text-lg font-semibold text-gray-700">Category</h3>
            <div className="h-auto overflow-y-auto rounded-lg p-3">
              {categories.map((cat) => (
                <div key={cat.id} className="mb-3">
                  <label className="flex items-center gap-2 font-medium cursor-pointer">
                    <input
                      type="checkbox"
                      checked={searchParams.get("category") === cat.slug}
                      onChange={() =>
                        handleFilterChange(
                          "category",
                          cat.slug === searchParams.get("category")
                            ? ""
                            : cat.slug
                        )
                      }
                      className="accent-primary"
                    />
                    {cat.name}
                  </label>
                </div>
              ))}
            </div>
          </div>

          {/* Price */}
          <div className="space-y-3 pt-6 border-t">
            <h3 className="text-lg font-semibold text-gray-700">Price</h3>
            <div className="flex gap-2">
              <input
                type="text"
                inputMode="numeric"
                value={minPriceInput}
                onChange={(e) =>
                  handlePriceInputChange("minPrice", e.target.value)
                }
                className="border rounded px-3 py-2 w-full text-sm"
                placeholder="Min"
              />
              <input
                type="text"
                inputMode="numeric"
                value={maxPriceInput}
                onChange={(e) =>
                  handlePriceInputChange("maxPrice", e.target.value)
                }
                className="border rounded px-3 py-2 w-full text-sm"
                placeholder="Max"
              />
            </div>
            <button
              onClick={applyPriceFilter}
              className="bg-primary hover:bg-primary/90 text-white text-sm rounded-md px-4 py-2 w-full"
            >
              Apply Filter
            </button>
          </div>

          {/* Rating */}
          <div className="space-y-3 pt-6 border-t">
            <h3 className="text-lg font-semibold text-gray-700">Rating</h3>
            <select
              className="border rounded px-3 py-2 w-full text-sm"
              value={searchParams.get("rating") || ""}
              onChange={(e) => handleFilterChange("rating", e.target.value)}
            >
              <option value="">All Ratings</option>
              {[4, 3, 2, 1].map((r) => (
                <option key={r} value={r}>
                  {r} and above
                </option>
              ))}
            </select>
          </div>
        </div>

        {/* Content */}
        <div className="col-span-4 md:col-span-3">
          <div className="mb-4">
            <Tabs defaultValue="view1">
              <TabsList className="flex items-center justify-between mb-4">
                <div>
                  <TabsTrigger value="view1">
                    <Grid2X2 />
                  </TabsTrigger>
                  <TabsTrigger value="view2">
                    <List />
                  </TabsTrigger>
                </div>
                <div>
                  <Select
                    onValueChange={(val) => handleFilterChange("sort", val)}
                  >
                    <SelectTrigger className="w-60">
                      <SelectValue placeholder="Sort by" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectLabel>Sort by</SelectLabel>
                        <SelectItem value="price_asc">Lower Price</SelectItem>
                        <SelectItem value="price_desc">Higher Price</SelectItem>
                        <SelectItem value="created_at_asc">Newest</SelectItem>
                        <SelectItem value="created_at_desc">Oldest</SelectItem>
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                </div>
              </TabsList>

              {/* Active filters */}
              {activeFilters.length > 0 && (
                <div className="flex flex-wrap items-center gap-2 mb-2 py-2 border-b">
                  {activeFilters.map(([key, val]) => (
                    <span
                      key={key}
                      className="bg-muted text-sm px-3 py-1 rounded-full flex items-center gap-1"
                    >
                      {key === "minPrice"
                        ? "Minimum Price"
                        : key === "maxPrice"
                        ? "Maximum Price"
                        : key === "rating"
                        ? `Rating ${val}+`
                        : val}
                      <button onClick={() => removeFilter(key)}>
                        <X size={14} />
                      </button>
                    </span>
                  ))}
                  <button
                    onClick={clearAllFilters}
                    className="text-sm text-green-600 hover:underline"
                  >
                    Clear All
                  </button>
                </div>
              )}

              {/* Product grid/list */}
              {data.length === 0 ? (
                <NoProductResult />
              ) : (
                <>
                  <TabsContent
                    className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6"
                    value="view1"
                  >
                    {data.map((product) => (
                      <ProductCard key={product.id} product={product} />
                    ))}
                  </TabsContent>
                  <TabsContent className="space-y-4" value="view2">
                    {data.map((product) => (
                      <div key={product.id}>
                        <ProductList product={product} />
                      </div>
                    ))}
                  </TabsContent>

                  {pagination && (
                    <Pagination
                      page={pagination.page}
                      limit={pagination.limit}
                      total={pagination.totalRows}
                      onPageChange={handlePageChange}
                    />
                  )}
                </>
              )}
            </Tabs>
          </div>
        </div>
      </div>
    </section>
  );
};

export default ProductResults;

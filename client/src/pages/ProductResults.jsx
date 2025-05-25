import { useMemo, useState } from "react";
import { Grid2X2, List, X } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useSearchParams } from "react-router-dom";
import { Pagination } from "@/components/ui/pagination";
import { useCategoriesQuery } from "@/hooks/useCategory";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { InputFilter } from "@/components/ui/InputFilter";
import { useSearchProductsQuery } from "@/hooks/useProduct";
import { SelectFilter } from "@/components/ui/SelectFilter";
import { CheckboxFilter } from "@/components/ui/CheckboxFilter";
import { productSortOptions, ratingOptions } from "@/lib/constant";
import { ProductCard } from "@/components/product-results/ProductCard";
import { ProductList } from "@/components/product-results/ProductList";
import { NoProductResult } from "@/components/product-results/NoProductResult";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { ProductResultsSkeleton } from "@/components/loading/ProductResultSkeleton";

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
      sort: params.sort || "created_at_desc",
      page: parseInt(params.page || "1"),
    };
  }, [searchParams]);

  const {
    data: searchData,
    isLoading,
    isError,
    refetch,
  } = useSearchProductsQuery({
    q: queryParams.q,
    page: queryParams.page,
    limit: 8,
    sort: queryParams.sort,
    category: queryParams.category,
    minPrice: queryParams.minPrice,
    maxPrice: queryParams.maxPrice,
    rating: queryParams.rating,
  });

  const { data: categoriesData, isLoading: isCategoriesLoading } =
    useCategoriesQuery({ limit: 5 });

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

  if (isLoading || isCategoriesLoading) return <ProductResultsSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const data = searchData?.data || [];

  const pagination = searchData?.pagination || {};

  return (
    <section className="section py-16 md:py-28 space-y-8">
      {showPriceWarning && (
        <div className="bg-red-100 text-red-600 text-sm p-3 rounded border border-red-300">
          Maximum price must be greater than minimum price.
        </div>
      )}

      <div className="grid grid-cols-4 gap-4">
        {/* Sidebar */}
        {/* Sidebar */}
        <div className="col-span-4 md:col-span-1">
          <div className="sticky top-24 space-y-4 max-h-[calc(100vh-6rem)] overflow-y-auto pr-2">
            <h2>Filter</h2>

            {/* Category */}
            <div className="space-y-2">
              <CheckboxFilter
                title="Category"
                options={categories.map((cat) => ({
                  value: cat.slug,
                  label: cat.name,
                }))}
                selectedValue={searchParams.get("category")}
                onChange={(val) => handleFilterChange("category", val)}
              />
            </div>

            {/* Price */}
            <div className="space-y-3 pt-6 border-t">
              <h4>Price</h4>
              <div className="flex gap-2">
                <InputFilter
                  placeholder="Min"
                  label="Minimum Price"
                  value={minPriceInput}
                  onChange={(val) => handlePriceInputChange("minPrice", val)}
                />
                <InputFilter
                  placeholder="Max"
                  label="Maximum Price"
                  value={maxPriceInput}
                  onChange={(val) => handlePriceInputChange("maxPrice", val)}
                />
              </div>
              <Button className="w-full" onClick={applyPriceFilter}>
                Apply Filter
              </Button>
            </div>

            {/* Rating */}
            <div className="space-y-3 pt-6 border-t">
              <h4>Rating</h4>
              <SelectFilter
                options={ratingOptions}
                className="w-full h-10"
                placeholder="All Ratings"
                value={searchParams.get("rating") || ""}
                onChange={(val) => handleFilterChange("rating", val)}
              />
            </div>
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
                  <SelectFilter
                    className="h-10 w-44 mb-4"
                    value={queryParams.sort}
                    options={productSortOptions}
                    onChange={(val) => handleFilterChange("sort", val)}
                  />
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

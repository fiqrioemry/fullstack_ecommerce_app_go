import { useDebounce } from "@/hooks/useDebounce";
import { Pagination } from "@/components/ui/pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { useCategoriesQuery } from "@/hooks/useCategory";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { useCategoryStore } from "@/store/useCategoryStore";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { SectionTitle } from "@/components/header/SectionTitle";
import { RecordNotFound } from "@/components/ui/RecordNotFound";
import { AddCategory } from "@/components/admin/categories/AddCategory";
import { CategoryCard } from "@/components/admin/categories/CategoryCard";

const CategoriesList = () => {
  const { page, limit, q, sort, setPage, setQ, setSort } = useCategoryStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = useCategoriesQuery({
    q: debouncedQ,
    page,
    limit,
    sort,
  });

  const categories = data?.data || [];

  const pagination = data?.pagination || null;

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle
        title="Categories List"
        description="Manage all categories and monitor product activities."
      />

      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by categories"}
        />
        <AddCategory />
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <LoadingSearch className="mt-10" />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : categories.length === 0 ? (
            <RecordNotFound title="No Category Found" q={q} />
          ) : (
            <CategoryCard
              sort={sort}
              setSort={setSort}
              categories={categories}
            />
          )}

          {pagination && pagination.totalRows > 10 && (
            <Pagination
              page={pagination.page}
              onPageChange={setPage}
              limit={pagination.limit}
              total={pagination.totalRows}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default CategoriesList;

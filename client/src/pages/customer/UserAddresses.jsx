import { useDebounce } from "@/hooks/useDebounce";
import { Pagination } from "@/components/ui/pagination";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserAddressesQuery } from "@/hooks/useAddress";
import { SelectFilter } from "@/components/ui/SelectFilter";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { NoAddress } from "@/components/customer/address/NoAddress";
import { AddAddress } from "@/components/customer/address/AddAddress";
import { AddressCard } from "@/components/customer/address/AddressCard";

const UserAddresses = () => {
  const { page, limit, q, sort, setPage, status, setQ, setStatus } =
    useQueryStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = useUserAddressesQuery({
    q: debouncedQ,
    page,
    limit,
    sort,
    status,
  });

  const addresses = data?.data || [];

  const pagination = data?.pagination;

  return (
    <section className="min-h-[45vh] space-y-6">
      {/* 🔍 Filter Bar */}
      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by Address / Province / city"}
        />

        <div className="flex gap-2 items-center">
          <SelectFilter
            value={status}
            onChange={setStatus}
            options={paymentStatusOptions}
          />
          <AddAddress />
        </div>
      </div>

      {isLoading ? (
        <LoadingSearch />
      ) : isError ? (
        <ErrorDialog onRetry={refetch} />
      ) : addresses.length === 0 ? (
        <NoAddress />
      ) : (
        <div className="space-y-4">
          {addresses.map((addr) => (
            <AddressCard key={addr.id} address={addr} />
          ))}
        </div>
      )}
      {pagination && (
        <Pagination
          page={pagination.page}
          onPageChange={setPage}
          limit={pagination.limit}
          total={pagination.totalRows}
        />
      )}
    </section>
  );
};

export default UserAddresses;

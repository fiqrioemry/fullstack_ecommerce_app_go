import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserAddressesQuery } from "@/hooks/useAddress";
import AddAddress from "@/components/customer/address/AddAddress";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import { NoAddress } from "@/components/customer/address/NoAddress";
import { AddressCard } from "@/components/customer/address/AddressCard";

const UserAddresses = () => {
  const { search, sort, page, limit, setSearch, setSort, setPage } =
    useQueryParamsStore();

  const { data, isLoading, isError, error } = useUserAddressesQuery(
    search,
    page,
    limit,
    sort
  );

  const addresses = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="min-h-[45vh] space-y-6">
      {/* 🔍 Filter Bar */}
      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
        <Input
          type="text"
          value={search}
          className="md:w-1/2"
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          placeholder="Cari nama / alamat / kota"
        />

        <div className="flex items-center gap-4">
          <select
            className="border rounded px-3 py-2 text-sm"
            value={sort}
            onChange={(e) => {
              setPage(1);
              setSort(e.target.value);
            }}
          >
            <option value="created_at desc">Terbaru</option>
            <option value="created_at asc">Terlama</option>
            <option value="name asc">Nama A-Z</option>
            <option value="name desc">Nama Z-A</option>
          </select>
          <AddAddress />
        </div>
      </div>

      {/* 📦 Content */}
      {isLoading ? (
        <Loading className="mt-10" />
      ) : isError ? (
        <ErrorDialog message={error?.message || "Terjadi kesalahan"} />
      ) : addresses.length === 0 ? (
        <NoAddress />
      ) : (
        <div className="space-y-4">
          {addresses.map((addr) => (
            <AddressCard key={addr.id} address={addr} />
          ))}

          {/* 📄 Pagination */}
          {pagination && (
            <div className="flex items-center justify-between pt-6">
              <Button
                variant="outline"
                onClick={() => setPage((p) => Math.max(p - 1, 1))}
                disabled={page === 1}
              >
                Sebelumnya
              </Button>
              <p className="text-sm text-muted-foreground">
                Page {pagination.page} / {pagination.totalPages}
              </p>
              <Button
                variant="outline"
                onClick={() => setPage((p) => p + 1)}
                disabled={page >= pagination.totalPages}
              >
                Selanjutnya
              </Button>
            </div>
          )}
        </div>
      )}
    </section>
  );
};

export default UserAddresses;

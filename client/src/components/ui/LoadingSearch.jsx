import { Loader } from "lucide-react";

const LoadingSearch = () => {
  return (
    <div className="min-h-[45vh] w-full flex flex-col items-center justify-center bg-gray-50 animate-fadeIn">
      <div className="h-full flex flex-col items-center space-y-4">
        <Loader size={48} className="animate-spin text-primary" />
        <p className="text-gray-600 text-sm tracking-wide">Searching ....</p>
      </div>
    </div>
  );
};

export { LoadingSearch };

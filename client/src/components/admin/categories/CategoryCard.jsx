import {
  Table,
  TableRow,
  TableBody,
  TableHead,
  TableCell,
  TableHeader,
} from "@/components/ui/table";
import { UpdateCategory } from "./UpdateCategory";
import { DeleteCategory } from "./DeleteCategory";
import { ChevronDown, ChevronUp } from "lucide-react";

export const CategoryCard = ({ categories, sort, setSort }) => {
  const renderSortIcon = (field) => {
    if (sort === `${field}_asc`)
      return <ChevronUp className="w-4 h-4 inline" />;
    if (sort === `${field}_desc`)
      return <ChevronDown className="w-4 h-4 inline" />;
    return null;
  };
  return (
    <div className="hidden md:block w-full">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="text-center">Preview</TableHead>
            <TableHead
              className="cursor-pointer"
              onClick={() => setSort("name")}
            >
              Name {renderSortIcon("name")}
            </TableHead>
            <TableHead className="text-center">Slug</TableHead>
            <TableHead className="text-center">Action</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody className="h-12">
          {categories.map((category) => (
            <TableRow key={category.id}>
              <TableCell className="text-center">
                <div className="flex items-center justify-center gap-4">
                  <img
                    src={category.image}
                    alt={category.name}
                    className="h-28 w-28 object-cover rounded"
                  />
                </div>
              </TableCell>
              <TableCell className="text-center">{category.name}</TableCell>
              <TableCell className="text-center">{category.slug}</TableCell>
              <TableCell className="text-center space-x-4">
                <UpdateCategory category={category} />
                <DeleteCategory category={category} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};

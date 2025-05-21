import { UpdateCategory } from "./UpdateCategory";
import { DeleteCategory } from "./DeleteCategory";
import { TableRow, TableCell } from "@/components/ui/table";

export const CategoryCard = ({ category }) => {
  return (
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
  );
};

import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import { format } from "date-fns";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { Progress } from "@/components/ui/progress";
import { Textarea } from "@/components/ui/textarea";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Checkbox } from "@/components/ui/checkbox";

import { cn } from "@/lib/utils";

import { ArrowLeftIcon, CalendarIcon, ImageUp } from "lucide-react";
import { z } from "zod";

import { useState } from "react";

const committees: [string] = ["QMISK", "TMEIT", "ITK", "SMN"];

const ReceiptReportSchema = z.object({
  full_name: z.string(),
  food_sum: z.number().nonnegative(),
  beer_sum: z.number().nonnegative(),
  soda_sum: z.number().nonnegative(),
  cider_sum: z.number().nonnegative(),
  wine_sum: z.number().nonnegative(),
  spirits_sum: z.number().nonnegative(),
  internrep_sum: z.number().nonnegative(),
  committee: z.enum(committees),
});

function ReceiptReport() {
  const [date, setDate] = useState<Date>();

  const onImageClick = () => {
    alert(1);
  };

  const pages = [
    <>
      {/* == Picture upload page == */}
      <h1 className="pt-4 ml-2 mb-3 text-xl">Upload a receipt picture</h1>
      <p className="ml-2 mb-3">The receipt should be clearly readable.</p>
      <div
        className="flex items-center justify-center m-3 grow bg-card border-1 rounded-lg mb-6"
        onClick={onImageClick}
      >
        <ImageUp className="h-16 w-16" />
      </div>
    </>,

    <>
      {/* == Readability Page == */}
      <h1 className="pt-4 ml-2 mb-1 text-xl">Is this clearly readable?</h1>
      <p className="ml-2 mb-1 text-sm">
        Make sure it's clearly readable, as you're responsible until cashier has accepted your report.
      </p>
      <div className="flex items-center justify-center m-3 grow rounded-lg mb-3" onClick={onImageClick}>
        <img src="receipt.jpg" className="rounded-lg object-contain" />
      </div>
      <div className="flex items-center pl-3 pr-3 mb-4 gap-2">
        <Button className="grow flex-1/2" variant="secondary">
          Continue
        </Button>
        <Button className="grow flex-1/2" variant="destructive">
          Retake the photo
        </Button>
      </div>
    </>,

    <>
      {/* == Purchase Contents == */}
      <div className="grid gap-2">
        <h1 className="pt-4 ml-2 text-xl">What has been bought?</h1>
        <p className="ml-2 text-sm">Document what has been bought.</p>
        <div className="mt-4 ml-3 mr-3">
          <Textarea className="h-[20dvh]"></Textarea>
        </div>
      </div>
    </>,

    <>
      <div className="grid gap-2">
        {/* == Select Purchase Date and Committee == */}
        <h1 className="pt-4 ml-2 mb-1 text-xl">Who and when?</h1>
        <p className="ml-2 mb-1 text-sm">
          Make sure it's clearly readable, as you're responsible until cashier has accepted your report.
        </p>

        <div className="content-center justify-center grid gap-4">
          <div>
            <h2>Which committee made this purchase?</h2>
            <Select>
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="Committee" />
              </SelectTrigger>
              <SelectContent>
                {committees.map((committee) => (
                  <SelectItem value={committee}>{committee}</SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
          <div>
            <h2>When was the purchase done?</h2>
            <Popover>
              <PopoverTrigger asChild>
                <Button
                  variant={"outline"}
                  className={cn(
                    "w-[280px] justify-start text-left font-normal",
                    !date && "text-muted-foreground",
                  )}
                >
                  <CalendarIcon className="mr-2 h-4 w-4" />
                  {date ? format(date, "PPP") : <span>Pick a date</span>}
                </Button>
              </PopoverTrigger>
              <PopoverContent className="w-auto p-0">
                <Calendar mode="single" selected={date} onSelect={setDate} initialFocus />
              </PopoverContent>
            </Popover>
          </div>
          <div className="grid gap-2">
            <h2>Who made this purchase?</h2>
            <Input disabled placeholder="Julle Juliusson Keuschnig" />
            <div className="items-top flex space-x-2">
              <Checkbox />
              <label
                htmlFor="terms1"
                className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              >
                Not you?
              </label>
            </div>
          </div>
        </div>
      </div>
    </>,
  ];

  const [index, setIndex] = useState(0);

  const nextPage = () => {
    setIndex((index + 1) % pages.length);
  };

  return (
    <>
      <div className="h-[calc(100dvh-60px)] ml-4 mr-4 pt-10 mb-4 flex flex-col justify-center items-stretch">
        <div className="flex ml-3 items-center gap-2 justify-center">
          <Button size="icon" className="inline mr-2" variant="secondary">
            <ArrowLeftIcon className="m-auto" />
          </Button>
          <Progress value={(100 / pages.length) * (index + 1)} className="mr-4" />
        </div>
        <div className="flex grow flex-col">{pages[index]}</div>
      </div>
      <img
        src="flashcow.webp"
        className="fixed z-50 rotate-[270deg] h-16 w-16 right-0 top-16"
        onClick={nextPage}
      />
    </>
  );
}

export default ReceiptReport;

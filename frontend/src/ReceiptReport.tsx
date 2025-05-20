import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import { Progress } from "@/components/ui/progress";
import { ArrowLeftIcon, ImageUp, RotateCwIcon } from "lucide-react";

import { useState } from "react";

function ReceiptReport() {
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
      <div className="flex items-center pl-3 pr-3 mb-4">
        <Button className="h-9 mr-2 grow" variant="secondary">
          Continue
        </Button>
        <Button size="icon" variant="destructive">
          <RotateCwIcon />
        </Button>
      </div>
    </>,

    <>
      {/* == Select Purchase Date and Committee == */}
      <h1 className="pt-4 ml-2 mb-1 text-xl">Who and when?</h1>
      <p className="ml-2 mb-1 text-sm">
        Make sure it's clearly readable, as you're responsible until cashier has accepted your report.
      </p>
      <Calendar></Calendar>
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
        {pages[index]}
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

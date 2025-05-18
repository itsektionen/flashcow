"use client";

import { ItBolt } from "@/components/it-bolt";
import { ItChip } from "@/components/it-chip";
import { Drawer, DrawerContent, DrawerHeader, DrawerTitle, DrawerTrigger } from "@/components/ui/drawer";
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from "@/components/ui/navigation-menu";
import { cn } from "@/lib/utils";
import { CalendarIcon, FileTextIcon, GraduationCapIcon, HouseIcon, ReceiptIcon } from "lucide-react";
import { Link } from "react-router";
import React, { useEffect, useState } from "react";
import { ThemeToggle } from "@/components/theme-toggle";

const DesktopNavbar = () => {
  return (
    <header className="hidden sm:block sticky z-50 top-0 border-b bg-background/90 backdrop-blur">
      <div className="container mx-auto px-6 h-16 border-x flex items-center gap-4">
        <Link className="flex text-sm text-foreground items-center gap-2 font-medium mr-2" to="/">
          <ItChip primary="var(--primary)" />
        </Link>
        <NavigationMenu>
          <NavigationMenuList>
            <NavigationMenuItem>
              <NavigationMenuTrigger>Education</NavigationMenuTrigger>
            </NavigationMenuItem>
            <NavigationMenuItem>
              <NavigationMenuTrigger>Chapter</NavigationMenuTrigger>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>
        <div className="ml-auto">
          <ThemeToggle></ThemeToggle>
        </div>
      </div>
    </header>
  );
};

const MobileNavbar = () => {
  const pathname = "";
  const [open, setOpen] = useState<boolean>(false);

  useEffect(() => {
    setOpen(false);
  }, [pathname]);

  return (
    <nav className="fixed flex sm:hidden items-center justify-between bottom-0 left-0 bg-card/90 backdrop-blur z-50 right-0 px-6 py-4 h-15 border-t">
      <MobileNavItem href="/">
        <HouseIcon />
      </MobileNavItem>
      <MobileNavItem href="/receiptreport">
        <ReceiptIcon />
      </MobileNavItem>
      <Drawer preventScrollRestoration open={open} onOpenChange={setOpen}>
        <DrawerTrigger className="active:scale-90 p-3 transition-transform h-10 w-10 bg-primary rounded-full drop-shadow text-primary-foreground flex items-center justify-center">
          <ItBolt primary="var(--card)" />
        </DrawerTrigger>
        <DrawerContent className="pb-10">
          <DrawerHeader className="flex items-center gap-2">
            <ItChip primary="var(--primary)" />
            <DrawerTitle>The IT Chapter</DrawerTitle>
          </DrawerHeader>
          <div className="px-6">
            <Link to={"/"} className="font-medium mb-4 block w-fit">
              Home
            </Link>
            <Link to={"/contact"} className="font-medium mb-4 block w-fit">
              Contact
            </Link>
          </div>
        </DrawerContent>
      </Drawer>
      <MobileNavItem href="/education/courses">
        <GraduationCapIcon strokeWidth={1.8} className="size-6" />
      </MobileNavItem>
      <MobileNavItem href="/documents/protocols">
        <FileTextIcon />
      </MobileNavItem>
    </nav>
  );
};

const MobileNavItem = ({
  href,
  children,
  className,
}: {
  href: string;
  children: React.ReactNode;
  className?: string;
}) => {
  const pathname = "";
  const isActive = pathname === href;
  return (
    <Link
      className={cn(
        "[&_svg:not([class*='size-'])]:size-5 [&_svg]:active:scale-80 [&_svg]:transition-all",
        isActive ? "[&_svg]:text-primary" : "[&_svg]:text-muted-foreground",
        className,
      )}
      to={href}
    >
      {children}
    </Link>
  );
};

const Navbar = () => {
  return (
    <>
      <DesktopNavbar />
      <MobileNavbar />
    </>
  );
};

export { Navbar };

const ListItem = React.forwardRef<React.ComponentRef<"a">, React.ComponentPropsWithoutRef<"a">>(
  ({ className, title, children, ...props }, ref) => {
    return (
      <li>
        <NavigationMenuLink asChild>
          <a
            ref={ref}
            className={cn(
              "block select-none flex-col items-start space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
              className,
            )}
            {...props}
          >
            <div className="text-sm font-medium leading-none">{title}</div>
            <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">{children}</p>
          </a>
        </NavigationMenuLink>
      </li>
    );
  },
);

ListItem.displayName = "ListItem";

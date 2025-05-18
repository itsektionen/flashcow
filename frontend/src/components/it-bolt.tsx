/**
 * The chapter emblem.
 *
 * @prop primary - The primary color of the bolt.
 * @prop secondary - The secondary color of the bolt (defaults to primary).
 * @returns
 */
const ItBolt = ({
  primary = "var(--foreground)",
  secondary = primary,
  size = 36,
  className,
  style,
}: {
  primary?: string;
  secondary?: string;
  className?: string;
  size?: number;
  style?: React.CSSProperties;
}) => {
  return (
    <svg
      version="1.1"
      id="IT"
      xmlns="http://www.w3.org/2000/svg"
      x="0px"
      y="0px"
      height={size}
      width={size}
      viewBox="0 0 339.9 564.1"
      className={className}
      style={style}
    >
      <path
        id="I"
        fill={primary}
        d="M29.8,331.5l132.8-56.4c4.4-1.9,7.3-6.2,7.3-11V3c0-3.3-4.5-4.2-5.8-1.2L25.9,327.6
	C24.9,330.1,27.4,332.6,29.8,331.5L29.8,331.5z"
      />
      <path
        id="T"
        fill={secondary}
        d="M4.8,370.7l154.3-31.4c5.6-1.1,10.8,3.1,10.8,8.8v212.9c0,3.3,4.5,4.2,5.8,1.2l101.4-239
	c2.3-5.5,7.2-9.4,13-10.6l5.3-1.1c6.8-1.4,12.4-6,15.1-12.4l28.8-67.9c2.1-5-2.9-10-7.9-7.9L2.5,363.1C-2,365,0.1,371.7,4.8,370.7z"
      />
    </svg>
  );
};

export { ItBolt };

const PageLoading = () => {
  return (
    <div
      style={{
        display: "flex",
        width: "100%",
        height: "100%",
        justifyContent: "center",
        alignItems: "center",
        color: "#5B67EA",
      }}
    >
      <svg
        width="24"
        height="24"
        viewBox="0 0 48 48"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g clip-path="url(#clip0_1262_42688)">
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M36.2528 31.3147C39.0815 26.3803 38.7345 20.4914 35.8438 16.0155C35.3045 15.1804 35.5443 14.0662 36.3794 13.5268C37.2145 12.9875 38.3287 13.2273 38.868 14.0624C42.4604 19.625 42.9003 26.9573 39.376 33.1051C35.8517 39.2529 29.3021 42.5784 22.6866 42.2892C21.6935 42.2457 20.9235 41.4054 20.967 40.4123C21.0104 39.4191 21.8507 38.6492 22.8439 38.6926C28.167 38.9253 33.4241 36.249 36.2528 31.3147ZM26.7899 7.73007C26.7003 8.72014 25.8251 9.45012 24.835 9.36053C19.2852 8.85833 13.6958 11.5484 10.7466 16.6929C7.79742 21.8374 8.30046 28.02 11.5382 32.5555C12.1158 33.3646 11.9282 34.4887 11.1191 35.0663C10.31 35.6439 9.18582 35.4562 8.60822 34.6471C4.58332 29.0091 3.9489 21.3123 7.62339 14.9025C11.2979 8.49271 18.2604 5.15089 25.1595 5.77518C26.1495 5.86477 26.8795 6.74001 26.7899 7.73007Z"
            fill="#5B67EA"
          />
        </g>
        <defs>
          <clipPath id="clip0_1262_42688">
            <rect width="48" height="48" fill="white" />
          </clipPath>
        </defs>
        <animateTransform
          attributeName="transform"
          type="rotate"
          from="0"
          to="360"
          dur="1s"
          repeatCount="indefinite"
        />
      </svg>
      <div style={{ fontSize: "20px" }}> Loading.... </div>
    </div>
  );
};

export default PageLoading;

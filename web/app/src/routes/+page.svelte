<!-- YOU CAN DELETE EVERYTHING IN THIS PAGE -->

<div class="container h-full mx-auto flex justify-center items-center">
	<div class="space-y-10 text-center flex flex-col items-center">
		<h2 class="h2">Welcome to boxie.</h2>
		<!-- Animated Logo -->
		<figure>
			<section class="img-bg" />
			<svg width="265" height="265" viewBox="0 0 265 265" fill="none" xmlns="http://www.w3.org/2000/svg">
				<g filter="url(#filter0_i_3_7)">
					<path d="M0 56C0 36.3982 0 26.5972 3.81477 19.1103C7.17034 12.5247 12.5247 7.17034 19.1103 3.81477C26.5972 0 36.3982 0 56 0H209C228.602 0 238.403 0 245.89 3.81477C252.475 7.17034 257.83 12.5247 261.185 19.1103C265 26.5972 265 36.3982 265 56V209C265 228.602 265 238.403 261.185 245.89C257.83 252.475 252.475 257.83 245.89 261.185C238.403 265 228.602 265 209 265H56C36.3982 265 26.5972 265 19.1103 261.185C12.5247 257.83 7.17034 252.475 3.81477 245.89C0 238.403 0 228.602 0 209V56Z" fill="white"/>
				</g>
				<rect x="64" y="93" width="28" height="55" rx="14" fill="#1E1E1E"/>
				<rect x="174" y="93" width="28" height="55" rx="14" fill="#1E1E1E"/>
				<path d="M112 170V170C121.242 185.553 143.758 185.553 153 170V170" stroke="#1E1E1E" stroke-width="10" stroke-linecap="round"/>
				<defs>
					<filter id="filter0_i_3_7" x="0" y="0" width="265" height="265" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
						<feFlood flood-opacity="0" result="BackgroundImageFix"/>
						<feBlend mode="normal" in="SourceGraphic" in2="BackgroundImageFix" result="shape"/>
						<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
						<feOffset dy="6"/>
						<feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1"/>
						<feColorMatrix type="matrix" values="0 0 0 0 0.129412 0 0 0 0 0.305882 0 0 0 0 0.45098 0 0 0 1 0"/>
						<feBlend mode="normal" in2="shape" result="effect1_innerShadow_3_7"/>
					</filter>
				</defs>
			</svg>

		</figure>
		<!-- / -->
		<div class="flex justify-center space-x-2">
			<a
				class="btn variant-filled"
				href="/app"
			>
				Control Panel
			</a>
		</div>
		<div class="space-y-2">
			<p><a class="underline" href="#">How to create an account</a></p>
		</div>
	</div>
</div>

<style lang="postcss">
	figure {
		@apply flex relative flex-col;
	}
	figure svg,
	.img-bg {
		@apply w-64 h-64 md:w-80 md:h-80;
	}
	.img-bg {
		@apply absolute z-[-1] rounded-full blur-[50px] transition-all;
		animation: pulse 5s cubic-bezier(0, 0, 0, 0.5) infinite,
			glow 5s linear infinite;
	}
	@keyframes glow {
		0% {
			@apply bg-primary-400/50;
		}
		33% {
			@apply bg-secondary-400/50;
		}
		66% {
			@apply bg-tertiary-400/50;
		}
		100% {
			@apply bg-primary-400/50;
		}
	}
	@keyframes pulse {
		50% {
			transform: scale(1.5);
		}
	}
</style>

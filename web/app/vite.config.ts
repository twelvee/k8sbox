import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	resolve: {
		alias: {
			$img: "/src/lib/img",
		},
	},
	plugins: [sveltekit()]
});

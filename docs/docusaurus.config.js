// @ts-check
// `@type` JSDoc annotations allow editor autocompletion and type checking
// (when paired with `@ts-check`).
// There are various equivalent ways to declare your Docusaurus config.
// See: https://docusaurus.io/docs/api/docusaurus-config

import { lightCodeTheme, darkCodeTheme } from 'prism-react-renderer';

const organizationName = '2024-T0002-EC09-G04';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Grupo 4 - Orbit City',
  tagline: 'Planejamento de hiperconectividade para cidades inteligentes',
  favicon: 'icons/inteli_logo.png',

  // Set the production url of your site here
  url: `https://${organizationName}.github.io`,
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: `/${organizationName}/`,

  projectName: '2024-T0002-EC09-G04',
  organizationName: 'docs',
  trailingSlash: false,

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'br',
    locales: ['br'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
      ({
        docs: {
          sidebarPath: './sidebars.js',
          routeBasePath: '/'
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      }),
    ],
  ],

  themeConfig:
  /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
  ({
    // Replace with your project's social card
    image: 'img/m9-social-card.jpg',
    navbar: {
      title: 'Grupo4 | Orbit City',
      logo: {
        alt: 'logo',
        src: 'icons/inteli_logo.png',
      },
      items: [
        {
          to: "https://github.com/Inteli-College/2024-T0002-EC09-G04",
          position: "right",
          className: "header-github-link",
          "aria-label": "GitHub repository",
        },
      ],
    },
    footer: {
      style: 'dark',
      copyright: `Copyright © ${new Date().getFullYear()} Grupo 4 | Orbit City`,
    },
    prism: {
      theme: lightCodeTheme,
      darkTheme: darkCodeTheme,
    },
    markdown: {
      mermaid: true,
    },    
  }),
};

export default config;

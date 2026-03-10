# cdn

<a href="https://gitpod.io/#https://github.com/gouniverse/cdn" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>
<a href="https://app.codeanywhere.com/#https://github.com/gouniverse/cdn/blob/292d9b4a74820f4526313285b852aaf4ec6390a5/README.md" style="float:right:"><img src="https://codeanywhere.com/img/open-in-codeanywhere-btn.svg" alt="Open in Gitpod" loading="lazy"></a>

![tests](https://github.com/gouniverse/cdn/workflows/tests/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/cdn)](https://goreportcard.com/report/github.com/gouniverse/cdn)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/cdn)](https://pkg.go.dev/github.com/gouniverse/cdn)

Easy to use CDN URLs for using in Golang web UIs

If you want faster loading, reliability, and total security consider using UNCDN (https://github.com/gouniverse/uncdn)

## Customizing CDN URL Prefix

The `CDN_URL_PREFIX` feature allows you to specify a custom CDN URL prefix via an environment variable. This can be useful if you want to use a different CDN for serving your assets.

To use the `CDN_URL_PREFIX` feature, set the `CDN_URL_PREFIX` environment variable to your desired CDN URL prefix. For example:

```sh
export CDN_URL_PREFIX="https://mycustomcdn.com/"
```

If the `CDN_URL_PREFIX` environment variable is set, it will be used as the base URL for the CDN; otherwise, the original CDN URL will be used.

## Available methods

### AlpineJS

- AlpineJs_3_15_8() - latest
- AlpineJs_3_13_8()
- AlpineJs_3_13_0()
- AlpineJs_3_12_3()

### AnimatedCSS

- AnimatedCSS_4_1_1 - latest

### Bootstrap

- BootstrapCss_5_3_8() - latest
- BootstrapJs_5_3_8()
- BootstrapCosmoCss_5_3_8() // + all the Bottswatch themes for 5.3.8

- BootstrapCss_5_3_7()
- BootstrapJs_5_3_7()
- BootstrapCosmoCss_5_3_7() // + all the Bottswatch themes for 5.3.7

- BootstrapCss_5_3_3()
- BootstrapJs_5_3_3()
- BootstrapCosmoCss_5_3_3() // + all the Bottswatch themes for 5.3.3

- BootstrapCss_5_3_2()
- BootstrapJs_5_3_2()
- BootstrapCosmoCss_5_3_2() // + all the Bottswatch themes for 5.3.2

- BootstrapCss_5_3_1()
- BootstrapJs_5_3_1()
- BootstrapCosmoCss_5_3_1() // + all the Bottswatch themes for 5.3.1

- BootstrapCss_5_3_0()
- BootstrapJs_5_3_0()
- BootstrapCosmoCss_5_3_0() // + all the Bottswatch themes for 5.3.0

- BootstrapCss_5_2_3()
- BootstrapJs_5_2_3()

### Bootstrap Icons

- BootstrapIconsCss_1_11_3() - latest
- BootstrapIconsCss_1_11_2()
- BootstrapIconsCss_1_11_0()
- BootstrapIconsCss_1_10_5()
- BootstrapIconsCss_1_10_2()
- BootstrapIconsCss_1_9_1()

- ChartsCss_0_9_3()

- FontAwesomeCss_7_2_0() - latest
- FontAwesomeCss_6_5_2()
- FontAwesomeCss_6_4_2()
- FontAwesomeCss_6_1_2()

### GoogleFont

- GoogleFont(family string, weight string)

Example:

```golang
cdn.GoogleFont("Josefin Sans", "300,400,500")
```

### HTMX

- Htmx_2_0_8() - latest
- Htmx_2_0_0()
- Htmx_1_9_11()
- Htmx_1_9_9()
- Htmx_1_9_6()
- Htmx_1_9_5()
- Htmx_1_9_4()
- Htmx_1_9_2()

### JQuery

- Jquery_4_0_0() - latest
- Jquery_3_7_1()
- Jquery_3_6_4()

### JQuery DataTables

- JqueryDataTablesCss_2_3_7() - latest
- JqueryDataTablesJs_2_3_7() - latest
- JqueryDataTablesCss_1_13_4()
- JqueryDataTablesJs_1_13_4()

### JQTree

- JqTreeCss_1_8_0()
- JqTreeJs_1_8_0()
- JqTreeCss_1_7_0()
- JqTreeJs_1_7_0()

### Notify

- Notify_0_4_2()

### SweetAlert v2

- Sweetalert2_11_26_22() - latest
- Sweetalert2_11()
- Sweetalert2_10()

### Tailwind

- TailwindCss_4_2_1() - latest
- TailwindCss_3_4_4()
- TailwindCss_3_3_3()

### Trumbowyg

- TrumbowygCss_2_31_0() - latest
- TrumbowygJs_2_31_0() - latest
- TrumbowygCss_2_27_3()
- TrumbowygJs_2_27_3()

### VueJS

- VueJs_3_5_30() - latest
- VueJs_3()
- VueElementPlusCss_2_3_8()
- VueElementPlusJs_2_3_8()
- VueTrumbowyg_4_0_0()

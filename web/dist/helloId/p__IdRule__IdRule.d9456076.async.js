"use strict";(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[984],{29595:function(ve,B,e){e.r(B),e.d(B,{default:function(){return oe}});var W=e(15009),R=e.n(W),G=e(99289),H=e.n(G),N=e(97857),j=e.n(N),Y=e(5574),I=e.n(Y),J=e(51042),K=e(11774),Q=e(97245),U=e(37476),Z=e(31199),X=e(5966),d=e(1413),L=e(91),A=e(22270),s=e(67294),M=e(66758),b=e(19323),t=e(85893),w=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","showSearch","options"],q=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","options"],k=function(r,m){var h=r.fieldProps,i=r.children,F=r.params,f=r.proFieldProps,P=r.mode,p=r.valueEnum,x=r.request,S=r.showSearch,o=r.options,v=(0,L.Z)(r,w),y=(0,s.useContext)(M.Z);return(0,t.jsx)(b.Z,(0,d.Z)((0,d.Z)({valueEnum:(0,A.h)(p),request:x,params:F,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,d.Z)({options:o,mode:P,showSearch:S,getPopupContainer:y.getPopupContainer},h),ref:m,proFieldProps:f},v),{},{children:i}))},_=s.forwardRef(function(l,r){var m=l.fieldProps,h=l.children,i=l.params,F=l.proFieldProps,f=l.mode,P=l.valueEnum,p=l.request,x=l.options,S=(0,L.Z)(l,q),o=(0,d.Z)({options:x,mode:f||"multiple",labelInValue:!0,showSearch:!0,showArrow:!1,autoClearSearchValue:!0,optionLabelProp:"label"},m),v=(0,s.useContext)(M.Z);return(0,t.jsx)(b.Z,(0,d.Z)((0,d.Z)({valueEnum:(0,A.h)(P),request:p,params:i,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,d.Z)({getPopupContainer:v.getPopupContainer},o),ref:r,proFieldProps:F},S),{},{children:h}))}),ee=s.forwardRef(k),ae=_,E=ee;E.SearchSelect=ae,E.displayName="ProFormComponent";var $=E,te=e(40741),re=e(15867),le=e(45360),V=e(48457),ne=function(){var r=(0,s.useState)(!1),m=I()(r,2),h=m[0],i=m[1],F=(0,s.useState)("\u65B0\u589EID\u89C4\u5219"),f=I()(F,2),P=f[0],p=f[1],x=te.Z.useForm(),S=I()(x,1),o=S[0],v=(0,s.useRef)(),y=(0,s.useState)(!1),z=I()(y,2),ue=z[0],O=z[1],se=[{title:"ID\u6807\u8BC6",dataIndex:"Type"},{title:"\u5F53\u524D\u9012\u589E\u503C",dataIndex:"CurrentId",search:!1},{title:"\u524D\u7F00",dataIndex:"Prefix",search:!1},{title:"\u540E\u7F00",dataIndex:"Suffix",search:!1},{title:"\u9012\u589E\u503C\u6700\u5C0F\u957F\u5EA6",dataIndex:"MinLength",search:!1},{title:"\u64CD\u4F5C",search:!1,render:function(u){return(0,t.jsx)(t.Fragment,{children:(0,t.jsx)("a",{onClick:function(){p("\u7F16\u8F91ID\u89C4\u5219"),i(!0);var a=j()({},u);a.Prefix=a.Prefix.split(",").filter(function(c){return c!==""}),a.Suffix=a.Suffix.split(",").filter(function(c){return c!==""}),o.setFieldsValue(a)},children:"\u7F16\u8F91"})})}}],ie=function(){return[(0,t.jsx)(re.ZP,{type:"primary",icon:(0,t.jsx)(J.Z,{}),onClick:function(){o.resetFields(),o.setFieldsValue({CurrentId:0,MinLength:0}),p("\u65B0\u589EID\u89C4\u5219"),i(!0)},children:"\u6DFB\u52A0"},"add")]},de=function(){var g=H()(R()().mark(function u(n){var a;return R()().wrap(function(C){for(;;)switch(C.prev=C.next){case 0:return a={data:[],success:!0,total:0},n.Page=n.current||1,n.PageSize=n.pageSize||10,C.next=5,(0,V.P3)("/api/getIdRuleList",n).then(function(T){a.success=T.code===200,a.data=T.data.List||[],a.data=a.data.map(function(D){return D.key=D.Id,D}),a.total=T.data.Total||0});case 5:return C.abrupt("return",Promise.resolve(a));case 6:case"end":return C.stop()}},u)}));return function(n){return g.apply(this,arguments)}}();return(0,t.jsxs)(K._z,{style:{minHeight:window.innerHeight-150},children:[(0,t.jsx)(Q.Z,{columns:se,search:{labelWidth:120},request:de,toolBarRender:ie,actionRef:v}),(0,t.jsxs)(U.Y,{width:500,title:P,open:h,form:o,loading:ue,modalProps:{onCancel:function(){return i(!1)}},onFinish:function(u){var n,a;O(!0),(0,V.dv)("/api/saveIdRule",j()(j()({},u),{},{Prefix:(n=u.Prefix)===null||n===void 0?void 0:n.join(","),Suffix:(a=u.Suffix)===null||a===void 0?void 0:a.join(",")})).then(function(c){c.code===200&&(le.ZP.success(c.message),i(!1),o.resetFields(),v.current&&v.current.reload()),O(!1)})},children:[(0,t.jsx)(Z.Z,{name:"Id",hidden:!0}),(0,t.jsx)(X.Z,{rules:[{required:!0,message:"\u8BF7\u8F93\u5165"}],name:"Type",label:"ID\u6807\u8BC6",placeholder:"\u8BF7\u8F93\u5165"}),(0,t.jsx)(Z.Z,{rules:[{type:"number",min:0}],name:"CurrentId",label:"\u5F53\u524D\u9012\u589E\u503C",placeholder:"\u8BF7\u8F93\u5165",tooltip:"ID\u9012\u589E\u7684\u8D77\u59CB\u503C"}),(0,t.jsx)($,{name:"Prefix",mode:"tags",label:"\u524D\u7F00",placeholder:"\u8BF7\u8F93\u5165",options:[{value:"\u5E74",label:"\u5E74"},{value:"\u6708",label:"\u6708"},{value:"\u65E5",label:"\u65E5"},{value:"\u65F6",label:"\u65F6"},{value:"\u5206",label:"\u5206"},{value:"\u79D2",label:"\u79D2"}]}),(0,t.jsx)($,{name:"Suffix",mode:"tags",label:"\u540E\u7F00",placeholder:"\u8BF7\u8F93\u5165",options:[{value:"\u5E74",label:"\u5E74"},{value:"\u6708",label:"\u6708"},{value:"\u65E5",label:"\u65E5"},{value:"\u65F6",label:"\u65F6"},{value:"\u5206",label:"\u5206"},{value:"\u79D2",label:"\u79D2"}]}),(0,t.jsx)(Z.Z,{rules:[{type:"number",min:0}],name:"MinLength",label:"\u9012\u589E\u503C\u6700\u5C0F\u957F\u5EA6",placeholder:"\u8BF7\u8F93\u5165",tooltip:"\u8BBE\u4E3A0\u65F6\u4E3A\u4E0D\u9650\u5236\u751F\u6210\u7684ID\u7684\u957F\u5EA6"})]})]})},oe=ne}}]);

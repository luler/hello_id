"use strict";(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[984],{29595:function(fe,R,e){e.r(R),e.d(R,{default:function(){return se}});var G=e(15009),L=e.n(G),H=e(99289),N=e.n(H),Y=e(97857),Z=e.n(Y),J=e(5574),I=e.n(J),K=e(51042),Q=e(11774),U=e(97245),X=e(37476),j=e(31199),w=e(5966),f=e(1413),A=e(91),M=e(22270),i=e(67294),O=e(66758),$=e(19323),t=e(85893),k=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","showSearch","options"],q=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","options"],_=function(l,m){var F=l.fieldProps,d=l.children,x=l.params,p=l.proFieldProps,P=l.mode,h=l.valueEnum,S=l.request,g=l.showSearch,u=l.options,c=(0,A.Z)(l,k),T=(0,i.useContext)(O.Z);return(0,t.jsx)($.Z,(0,f.Z)((0,f.Z)({valueEnum:(0,M.h)(h),request:S,params:x,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,f.Z)({options:u,mode:P,showSearch:g,getPopupContainer:T.getPopupContainer},F),ref:m,proFieldProps:p},c),{},{children:d}))},ee=i.forwardRef(function(r,l){var m=r.fieldProps,F=r.children,d=r.params,x=r.proFieldProps,p=r.mode,P=r.valueEnum,h=r.request,S=r.options,g=(0,A.Z)(r,q),u=(0,f.Z)({options:S,mode:p||"multiple",labelInValue:!0,showSearch:!0,showArrow:!1,autoClearSearchValue:!0,optionLabelProp:"label"},m),c=(0,i.useContext)(O.Z);return(0,t.jsx)($.Z,(0,f.Z)((0,f.Z)({valueEnum:(0,M.h)(P),request:h,params:d,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,f.Z)({getPopupContainer:c.getPopupContainer},u),ref:l,proFieldProps:x},g),{},{children:F}))}),te=i.forwardRef(_),ae=ee,y=te;y.SearchSelect=ae,y.displayName="ProFormComponent";var b=y,le=e(40741),ne=e(96074),re=e(5914),V=e(45360),oe=e(15867),E=e(48457),ue=function(){var l=(0,i.useState)(!1),m=I()(l,2),F=m[0],d=m[1],x=(0,i.useState)("\u65B0\u589EID\u89C4\u5219"),p=I()(x,2),P=p[0],h=p[1],S=le.Z.useForm(),g=I()(S,1),u=g[0],c=(0,i.useRef)(),T=(0,i.useState)(!1),z=I()(T,2),ie=z[0],W=z[1],de=[{title:"ID\u6807\u8BC6",dataIndex:"Type"},{title:"\u5F53\u524D\u9012\u589E\u503C",dataIndex:"CurrentId",search:!1},{title:"\u524D\u7F00",dataIndex:"Prefix",search:!1},{title:"\u540E\u7F00",dataIndex:"Suffix",search:!1},{title:"\u9012\u589E\u503C\u6700\u5C0F\u957F\u5EA6",dataIndex:"MinLength",search:!1},{title:"\u64CD\u4F5C",search:!1,render:function(o){return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)("a",{onClick:function(){h("\u7F16\u8F91ID\u89C4\u5219"),d(!0);var a=Z()({},o);a.Prefix=a.Prefix.split(",").filter(function(s){return s!==""}),a.Suffix=a.Suffix.split(",").filter(function(s){return s!==""}),u.setFieldsValue(a)},children:"\u7F16\u8F91"}),(0,t.jsx)(ne.Z,{type:"vertical"}),(0,t.jsx)("a",{style:{color:"red"},onClick:function(){re.Z.confirm({title:"\u60A8\u786E\u5B9A\u8981\u5220\u9664\u5417\uFF1F",content:"\u5220\u9664\u540E\uFF0C\u6570\u636E\u5C06\u65E0\u6CD5\u6062\u590D\uFF0C\u8BF7\u614E\u91CD\uFF01",onOk:function(s){(0,E.dv)("/api/delIdRule",{Ids:[o.Id]}).then(function(v){v.code===200&&(V.ZP.success(v.message),c.current.reload(),s())})}})},children:"\u5220\u9664"})]})}}],ce=function(){return[(0,t.jsx)(oe.ZP,{type:"primary",icon:(0,t.jsx)(K.Z,{}),onClick:function(){u.resetFields(),u.setFieldsValue({CurrentId:0,MinLength:0}),h("\u65B0\u589EID\u89C4\u5219"),d(!0)},children:"\u6DFB\u52A0"},"add")]},ve=function(){var C=N()(L()().mark(function o(n){var a;return L()().wrap(function(v){for(;;)switch(v.prev=v.next){case 0:return a={data:[],success:!0,total:0},n.Page=n.current||1,n.PageSize=n.pageSize||10,v.next=5,(0,E.P3)("/api/getIdRuleList",n).then(function(D){a.success=D.code===200,a.data=D.data.List||[],a.data=a.data.map(function(B){return B.key=B.Id,B}),a.total=D.data.Total||0});case 5:return v.abrupt("return",Promise.resolve(a));case 6:case"end":return v.stop()}},o)}));return function(n){return C.apply(this,arguments)}}();return(0,t.jsxs)(Q._z,{style:{minHeight:window.innerHeight-150},children:[(0,t.jsx)(U.Z,{columns:de,search:{labelWidth:120},request:ve,toolBarRender:ce,actionRef:c}),(0,t.jsxs)(X.Y,{width:500,title:P,open:F,form:u,loading:ie,modalProps:{onCancel:function(){return d(!1)}},onFinish:function(o){var n,a;W(!0),(0,E.dv)("/api/saveIdRule",Z()(Z()({},o),{},{Prefix:(n=o.Prefix)===null||n===void 0?void 0:n.join(","),Suffix:(a=o.Suffix)===null||a===void 0?void 0:a.join(",")})).then(function(s){s.code===200&&(V.ZP.success(s.message),d(!1),u.resetFields(),c.current&&c.current.reload()),W(!1)})},children:[(0,t.jsx)(j.Z,{name:"Id",hidden:!0}),(0,t.jsx)(w.Z,{rules:[{required:!0,message:"\u8BF7\u8F93\u5165"}],name:"Type",label:"ID\u6807\u8BC6",placeholder:"\u8BF7\u8F93\u5165"}),(0,t.jsx)(j.Z,{rules:[{type:"number",min:0}],name:"CurrentId",label:"\u5F53\u524D\u9012\u589E\u503C",placeholder:"\u8BF7\u8F93\u5165",tooltip:"ID\u9012\u589E\u7684\u8D77\u59CB\u503C"}),(0,t.jsx)(b,{name:"Prefix",mode:"tags",label:"\u524D\u7F00",placeholder:"\u8BF7\u8F93\u5165",options:[{value:"\u5E74",label:"\u5E74"},{value:"\u6708",label:"\u6708"},{value:"\u65E5",label:"\u65E5"},{value:"\u65F6",label:"\u65F6"},{value:"\u5206",label:"\u5206"},{value:"\u79D2",label:"\u79D2"}]}),(0,t.jsx)(b,{name:"Suffix",mode:"tags",label:"\u540E\u7F00",placeholder:"\u8BF7\u8F93\u5165",options:[{value:"\u5E74",label:"\u5E74"},{value:"\u6708",label:"\u6708"},{value:"\u65E5",label:"\u65E5"},{value:"\u65F6",label:"\u65F6"},{value:"\u5206",label:"\u5206"},{value:"\u79D2",label:"\u79D2"}]}),(0,t.jsx)(j.Z,{rules:[{type:"number",min:0}],name:"MinLength",label:"\u9012\u589E\u503C\u6700\u5C0F\u957F\u5EA6",placeholder:"\u8BF7\u8F93\u5165",tooltip:"\u8BBE\u4E3A0\u65F6\u4E3A\u4E0D\u9650\u5236\u751F\u6210\u7684ID\u7684\u957F\u5EA6"})]})]})},se=ue}}]);
"use strict";(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[984],{29595:function(fe,T,e){e.r(T),e.d(T,{default:function(){return de}});var H=e(15009),R=e.n(H),N=e(99289),Y=e.n(N),J=e(97857),D=e.n(J),K=e(5574),S=e.n(K),Q=e(51042),U=e(11774),X=e(97245),w=e(37476),j=e(31199),k=e(5966),p=e(1413),L=e(91),b=e(22270),d=e(67294),$=e(66758),M=e(19323),t=e(85893),q=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","showSearch","options"],_=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","options"],ee=function(a,v){var F=a.fieldProps,c=a.children,h=a.params,m=a.proFieldProps,x=a.mode,f=a.valueEnum,C=a.request,P=a.showSearch,s=a.options,i=(0,L.Z)(a,q),y=(0,d.useContext)($.Z);return(0,t.jsx)(M.Z,(0,p.Z)((0,p.Z)({valueEnum:(0,b.h)(f),request:C,params:h,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,p.Z)({options:s,mode:x,showSearch:P,getPopupContainer:y.getPopupContainer},F),ref:v,proFieldProps:m},i),{},{children:c}))},te=d.forwardRef(function(o,a){var v=o.fieldProps,F=o.children,c=o.params,h=o.proFieldProps,m=o.mode,x=o.valueEnum,f=o.request,C=o.options,P=(0,L.Z)(o,_),s=(0,p.Z)({options:C,mode:m||"multiple",labelInValue:!0,showSearch:!0,showArrow:!1,autoClearSearchValue:!0,optionLabelProp:"label"},v),i=(0,d.useContext)($.Z);return(0,t.jsx)(M.Z,(0,p.Z)((0,p.Z)({valueEnum:(0,b.h)(x),request:f,params:c,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,p.Z)({getPopupContainer:i.getPopupContainer},s),ref:a,proFieldProps:h},P),{},{children:F}))}),ue=d.forwardRef(ee),ae=te,I=ue;I.SearchSelect=ae,I.displayName="ProFormComponent";var V=I,le=e(40741),ne=e(5914),oe=e(12096),z=e(96074),re=e(86738),O=e(45360),se=e(15867),Z=e(48457),ie=function(){var a=(0,d.useState)(!1),v=S()(a,2),F=v[0],c=v[1],h=(0,d.useState)("\u65B0\u589EID\u89C4\u5219"),m=S()(h,2),x=m[0],f=m[1],C=le.Z.useForm(),P=S()(C,1),s=P[0],i=(0,d.useRef)(),y=(0,d.useState)(!1),W=S()(y,2),ce=W[0],G=W[1],pe=[{title:"ID\u6807\u8BC6",dataIndex:"type"},{title:"\u5F53\u524D\u9012\u589E\u503C",dataIndex:"currentId",search:!1},{title:"\u9884\u5360\u5E45\u5EA6",dataIndex:"step",search:!1},{title:"\u524D\u7F00",dataIndex:"prefix",search:!1},{title:"\u540E\u7F00",dataIndex:"suffix",search:!1},{title:"\u9012\u589E\u503C\u6700\u5C0F\u957F\u5EA6",dataIndex:"minLength",search:!1},{title:"\u521B\u5EFA\u65F6\u95F4",dataIndex:"createdAt",search:!1},{title:"\u64CD\u4F5C",search:!1,render:function(n){return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)("a",{onClick:function(){(0,Z.P3)("/api/createId",{type:n.type}).then(function(u){u.code===200&&(ne.Z.info({closable:!0,icon:null,footer:null,title:'\u6807\u8BC6"'+n.type+'"\u6210\u529F\u751F\u6210\u5982\u4E0BID',content:(0,t.jsx)("div",{style:{paddingBottom:20,paddingTop:10},children:(0,t.jsx)(oe.Z,{value:u.data[0],disabled:!0})})}),setTimeout(function(){i.current.reload()},1500))})},children:"\u751F\u6210"}),(0,t.jsx)(z.Z,{type:"vertical"}),(0,t.jsx)("a",{onClick:function(){f("\u7F16\u8F91ID\u89C4\u5219"),c(!0);var u=D()({},n);u.prefix=u.prefix.split(",").filter(function(r){return r!==""}),u.suffix=u.suffix.split(",").filter(function(r){return r!==""}),s.setFieldsValue(u)},children:"\u7F16\u8F91"}),(0,t.jsx)(z.Z,{type:"vertical"}),(0,t.jsx)(re.Z,{title:"\u60A8\u786E\u5B9A\u8981\u5220\u9664\u5417\uFF1F",description:"\u5220\u9664\u540E\uFF0C\u6570\u636E\u5C06\u65E0\u6CD5\u6062\u590D\uFF0C\u8BF7\u614E\u91CD\uFF01",onConfirm:function(u){(0,Z.dv)("/api/delIdRule",{ids:[n.id]}).then(function(r){r.code===200&&(O.ZP.success(r.message),i.current.reload())})},children:(0,t.jsx)("a",{style:{color:"red"},children:"\u5220\u9664"})})]})}}],ve=function(){return[(0,t.jsx)(se.ZP,{type:"primary",icon:(0,t.jsx)(Q.Z,{}),onClick:function(){s.resetFields(),s.setFieldsValue({currentId:0,step:100,minLength:0}),f("\u65B0\u589EID\u89C4\u5219"),c(!0)},children:"\u6DFB\u52A0"},"add")]},me=function(){var g=Y()(R()().mark(function n(l){var u;return R()().wrap(function(E){for(;;)switch(E.prev=E.next){case 0:return u={data:[],success:!0,total:0},l.page=l.current||1,l.pageSize=l.pageSize||10,E.next=5,(0,Z.P3)("/api/getIdRuleList",l).then(function(B){u.success=B.code===200,u.data=B.data.list||[],u.data=u.data.map(function(A){return A.key=A.id,A}),u.total=B.data.total||0});case 5:return E.abrupt("return",Promise.resolve(u));case 6:case"end":return E.stop()}},n)}));return function(l){return g.apply(this,arguments)}}();return(0,t.jsxs)(U._z,{style:{minHeight:window.innerHeight-150},children:[(0,t.jsx)(X.Z,{columns:pe,search:{labelWidth:120},request:me,toolBarRender:ve,actionRef:i}),(0,t.jsxs)(w.Y,{width:500,title:x,open:F,form:s,loading:ce,modalProps:{onCancel:function(){return c(!1)}},onFinish:function(n){var l,u;G(!0),(0,Z.dv)("/api/saveIdRule",D()(D()({},n),{},{Prefix:(l=n.Prefix)===null||l===void 0?void 0:l.join(","),Suffix:(u=n.Suffix)===null||u===void 0?void 0:u.join(",")})).then(function(r){r.code===200&&(O.ZP.success(r.message),c(!1),s.resetFields(),i.current&&i.current.reload()),G(!1)})},children:[(0,t.jsx)(j.Z,{name:"id",hidden:!0}),(0,t.jsx)(k.Z,{rules:[{required:!0,message:"\u8BF7\u8F93\u5165"}],name:"type",label:"ID\u6807\u8BC6",placeholder:"\u8BF7\u8F93\u5165"}),(0,t.jsx)(j.Z,{rules:[{type:"number",min:0}],name:"currentId",label:"\u5F53\u524D\u9012\u589E\u503C",placeholder:"\u8BF7\u8F93\u5165",tooltip:"ID\u9012\u589E\u7684\u8D77\u59CB\u503C"}),(0,t.jsx)(j.Z,{rules:[{type:"number",min:1}],name:"step",label:"\u9884\u5360\u5E45\u5EA6",placeholder:"\u8BF7\u8F93\u5165",tooltip:"\u7A0B\u5E8F\u6BCF\u6B21\u5F00\u59CB\u5206\u914D\u9884\u5148\u5360\u7528ID\u7684\u957F\u5EA6\uFF0C\u5206\u914D\u5B8C\u540E\u7EE7\u7EED\u83B7\u53D6\uFF0C\u8FD9\u6709\u52A9\u4E8E\u63D0\u9AD8\u751F\u6210\u901F\u5EA6"}),(0,t.jsx)(V,{name:"prefix",mode:"tags",label:"\u524D\u7F00",placeholder:"\u8BF7\u8F93\u5165",options:[{value:"\u5E74",label:"\u5E74"},{value:"\u6708",label:"\u6708"},{value:"\u65E5",label:"\u65E5"},{value:"\u65F6",label:"\u65F6"},{value:"\u5206",label:"\u5206"},{value:"\u79D2",label:"\u79D2"}]}),(0,t.jsx)(V,{name:"suffix",mode:"tags",label:"\u540E\u7F00",placeholder:"\u8BF7\u8F93\u5165",options:[{value:"\u5E74",label:"\u5E74"},{value:"\u6708",label:"\u6708"},{value:"\u65E5",label:"\u65E5"},{value:"\u65F6",label:"\u65F6"},{value:"\u5206",label:"\u5206"},{value:"\u79D2",label:"\u79D2"}]}),(0,t.jsx)(j.Z,{rules:[{type:"number",min:0}],name:"minLength",label:"\u9012\u589E\u503C\u6700\u5C0F\u957F\u5EA6",placeholder:"\u8BF7\u8F93\u5165",tooltip:"\u8BBE\u4E3A0\u65F6\u4E3A\u4E0D\u9650\u5236\u751F\u6210\u7684ID\u7684\u957F\u5EA6"})]})]})},de=ie}}]);
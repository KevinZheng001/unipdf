//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package context ;import (_ee "errors";_a "github.com/golang/freetype/truetype";_f "github.com/unidoc/unipdf/v3/core";_c "github.com/unidoc/unipdf/v3/internal/textencoding";_dc "github.com/unidoc/unipdf/v3/internal/transform";_cg "github.com/unidoc/unipdf/v3/model";_g "golang.org/x/image/font";_d "image";_eg "image/color";);func (_cab *TextState )Reset (){_cab .Tm =_dc .IdentityMatrix ();_cab .Tlm =_dc .IdentityMatrix ()};func (_egc *TextState )ProcQ (data []byte ,ctx Context ){_egc .ProcTStar ();_egc .ProcTj (data ,ctx )};type Context interface{Push ();Pop ();Matrix ()_dc .Matrix ;SetMatrix (_af _dc .Matrix );Translate (_ga ,_ge float64 );Scale (_fc ,_gf float64 );Rotate (_afa float64 );MoveTo (_ad ,_gff float64 );LineTo (_b ,_eeb float64 );CubicTo (_ebf ,_bb ,_fg ,_cf ,_ege ,_bc float64 );QuadraticTo (_gg ,_ae ,_df ,_eec float64 );NewSubPath ();ClosePath ();ClearPath ();Clip ();ClipPreserve ();ResetClip ();LineWidth ()float64 ;SetLineWidth (_cc float64 );SetLineCap (_aa LineCap );SetLineJoin (_db LineJoin );SetDash (_fb ...float64 );SetDashOffset (_ac float64 );Fill ();FillPreserve ();Stroke ();StrokePreserve ();SetRGBA (_gac ,_gb ,_ec ,_ab float64 );SetFillRGBA (_ef ,_ff ,_dfc ,_gc float64 );SetFillStyle (_acb Pattern );SetFillRule (_aca FillRule );SetStrokeRGBA (_ce ,_de ,_fa ,_cae float64 );SetStrokeStyle (_fgg Pattern );TextState ()*TextState ;DrawString (_eef string ,_dcc ,_ggb float64 );MeasureString (_fbf string )(_ggc ,_aef float64 );DrawRectangle (_fcf ,_agf ,_gbb ,_faf float64 );DrawImage (_gbf _d .Image ,_bcf ,_aac int );DrawImageAnchored (_egd _d .Image ,_dbb ,_afc int ,_cfd ,_gca float64 );Height ()int ;Width ()int ;};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;);func NewTextState ()*TextState {return &TextState {Th :100,Tm :_dc .IdentityMatrix (),Tlm :_dc .IdentityMatrix ()};};type FillRule int ;func NewTextFontFromPath (filePath string ,size float64 )(*TextFont ,error ){_bcc ,_gffe :=_cg .NewPdfFontFromTTFFile (filePath );if _gffe !=nil {return nil ,_gffe ;};return NewTextFont (_bcc ,size );};type Pattern interface{ColorAt (_dg ,_eb int )_eg .Color ;};func (_gd *TextFont )WithSize (size float64 ,originalFont *_cg .PdfFont )*TextFont {if size <=1{size =10;};return &TextFont {Font :_gd .Font ,Face :_a .NewFace (_gd ._gfc ,&_a .Options {Size :size }),Size :size ,_gfc :_gd ._gfc ,_fgc :originalFont };};func (_deb *TextFont )CharcodesToUnicode (charcodes []_c .CharCode )[]rune {if _deb ._fgc !=nil {return _deb ._fgc .CharcodesToUnicode (charcodes );};return _deb .Font .CharcodesToUnicode (charcodes );};func (_gfa *TextState )ProcTf (font *TextFont ){_gfa .Tf =font };type LineCap int ;func (_fe *TextState )ProcTd (tx ,ty float64 ){_fe .Tlm .Concat (_dc .TranslationMatrix (tx ,-ty ));_fe .Tm =_fe .Tlm .Clone ();};func (_dfb *TextState )ProcTD (tx ,ty float64 ){_dfb .Tl =-ty ;_dfb .ProcTd (tx ,ty )};type TextState struct{Tc float64 ;Tw float64 ;Th float64 ;Tl float64 ;Tf *TextFont ;Ts float64 ;Tm _dc .Matrix ;Tlm _dc .Matrix ;};func (_bbb *TextFont )GetRuneMetrics (r rune )(float64 ,float64 ,bool ){if _gad ,_age :=_bbb .Font .GetRuneMetrics (r );_age &&_gad .Wx !=0{return _gad .Wx ,_gad .Wy ,_age ;};if _bbb ._fgc ==nil {return 0,0,false ;};_aced ,_gae :=_bbb ._fgc .GetRuneMetrics (r );return _aced .Wx ,_aced .Wy ,_gae &&_aced .Wx !=0;};type LineJoin int ;func (_gfe *TextState )ProcTm (a ,b ,c ,d ,e ,f float64 ){_gfe .Tm =_dc .NewMatrix (a ,b ,c ,d ,e ,-f );_gfe .Tlm =_gfe .Tm .Clone ();};func NewTextFont (font *_cg .PdfFont ,size float64 )(*TextFont ,error ){_agff :=font .FontDescriptor ();if _agff ==nil {return nil ,_ee .New ("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072");};_efa ,_ea :=_f .GetStream (_agff .FontFile2 );if !_ea {return nil ,_ee .New ("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d");};_dfd ,_cfe :=_f .DecodeStream (_efa );if _cfe !=nil {return nil ,_cfe ;};_cd ,_cfe :=_a .Parse (_dfd );if _cfe !=nil {return nil ,_cfe ;};if size <=1{size =10;};return &TextFont {Font :font ,Face :_a .NewFace (_cd ,&_a .Options {Size :size }),Size :size ,_gfc :_cd },nil ;};const (LineCapRound LineCap =iota ;LineCapButt ;LineCapSquare ;);func (_dgcg *TextState )ProcTStar (){_dgcg .ProcTd (0,-_dgcg .Tl )};type TextFont struct{Font *_cg .PdfFont ;Face _g .Face ;Size float64 ;_gfc *_a .Font ;_fgc *_cg .PdfFont ;};type Gradient interface{Pattern ;AddColorStop (_ag float64 ,_ca _eg .Color );};func (_fbc *TextState )ProcDQ (data []byte ,aw ,ac float64 ,ctx Context ){_fbc .Tw =aw ;_fbc .Tc =ac ;_fbc .ProcQ (data ,ctx );};func (_dgc *TextFont )GetCharMetrics (code _c .CharCode )(float64 ,float64 ,bool ){if _bd ,_dgf :=_dgc .Font .GetCharMetrics (code );_dgf &&_bd .Wx !=0{return _bd .Wx ,_bd .Wy ,_dgf ;};if _dgc ._fgc ==nil {return 0,0,false ;};_gacd ,_fd :=_dgc ._fgc .GetCharMetrics (code );return _gacd .Wx ,_gacd .Wy ,_fd &&_gacd .Wx !=0;};func (_feb *TextState )Translate (tx ,ty float64 ){_feb .Tm =_dc .TranslationMatrix (tx ,ty ).Mult (_feb .Tm );};func (_gge *TextState )ProcTj (data []byte ,ctx Context ){_cda :=_gge .Tf .Size ;_ffe :=_gge .Th /100.0;_gcd :=_dc .NewMatrix (_cda *_ffe ,0,0,_cda ,0,_gge .Ts );_aee :=_gge .Tf .CharcodesToUnicode (_gge .Tf .BytesToCharcodes (data ));for _ ,_ece :=range _aee {if _ece =='\x00'{continue ;};_ggg :=_gge .Tm .Clone ();_gge .Tm .Concat (_gcd );_aab ,_ba :=_gge .Tm .Transform (0,0);ctx .Scale (1,-1);ctx .DrawString (string (_ece ),_aab ,_ba );ctx .Scale (1,-1);_ded :=0.0;if _ece ==' '{_ded =_gge .Tw ;};var _gbfc float64 ;if _ebe ,_ ,_efad :=_gge .Tf .GetRuneMetrics (_ece );_efad {_gbfc =_ebe *0.001*_cda ;}else {_gbfc ,_ =ctx .MeasureString (string (_ece ));};_bf :=(_gbfc +_gge .Tc +_ded )*_ffe ;_gge .Tm =_dc .TranslationMatrix (_bf ,0).Mult (_ggg );};};const (FillRuleWinding FillRule =iota ;FillRuleEvenOdd ;);func (_ace *TextFont )BytesToCharcodes (data []byte )[]_c .CharCode {if _ace ._fgc !=nil {return _ace ._fgc .BytesToCharcodes (data );};return _ace .Font .BytesToCharcodes (data );};
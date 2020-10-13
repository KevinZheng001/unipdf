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

package testutils ;import (_c "crypto/md5";_gd "encoding/hex";_fa "errors";_d "fmt";_cc "github.com/unidoc/unipdf/v3/common";_bd "github.com/unidoc/unipdf/v3/core";_adc "image";_ad "image/png";_g "io";_ab "os";_b "os/exec";_aa "path/filepath";_a "strings";_fe "testing";);func RunRenderTest (t *_fe .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ){_efde :=_a .TrimSuffix (_aa .Base (pdfPath ),_aa .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_efg *_fe .T ){_dag :=_aa .Join (outputDir ,_efde );_ddb :=_dag +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _ff :=RenderPDFToPNGs (pdfPath ,0,_ddb );_ff !=nil {_efg .Skip (_ff );};for _cad :=1;true ;_cad ++{_fgc :=_d .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_dag ,_cad );_cfa :=_aa .Join (baselineRenderPath ,_d .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_efde ,_cad ));if _ ,_baf :=_ab .Stat (_fgc );_baf !=nil {break ;};_efg .Logf ("\u0025\u0073",_cfa );if _ ,_ac :=_ab .Stat (_cfa );_ab .IsNotExist (_ac ){if saveBaseline {_efg .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_fgc ,_cfa );CopyFile (_fgc ,_cfa );continue ;};break ;};_efg .Run (_d .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_cad ),func (_af *_fe .T ){_af .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_fgc ,_cfa );_fbd ,_gg :=ComparePNGFiles (_fgc ,_cfa );if _ab .IsNotExist (_gg ){_af .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_fbd {_af .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");};});};});};func CompareDictionariesDeep (d1 ,d2 *_bd .PdfObjectDictionary )bool {if len (d1 .Keys ())!=len (d2 .Keys ()){_cc .Log .Debug ("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",len (d1 .Keys ()),len (d2 .Keys ()));_cc .Log .Debug ("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027",d1 .WriteString (),d2 .WriteString ());return false ;};for _ ,_ddg :=range d1 .Keys (){if _ddg =="\u0050\u0061\u0072\u0065\u006e\u0074"{continue ;};_edd :=_bd .TraceToDirectObject (d1 .Get (_ddg ));_daf :=_bd .TraceToDirectObject (d2 .Get (_ddg ));if _edd ==nil {_cc .Log .Debug ("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l");return false ;};if _daf ==nil {_cc .Log .Debug ("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l");return false ;};switch _ec :=_edd .(type ){case *_bd .PdfObjectDictionary :_egf ,_ccd :=_daf .(*_bd .PdfObjectDictionary );if !_ccd {_cc .Log .Debug ("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054",_edd ,_daf );return false ;};if !CompareDictionariesDeep (_ec ,_egf ){return false ;};continue ;case *_bd .PdfObjectArray :_bage ,_aaa :=_daf .(*_bd .PdfObjectArray );if !_aaa {_cc .Log .Debug ("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");return false ;};if _ec .Len ()!=_bage .Len (){_cc .Log .Debug ("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",_ec .Len (),_bage .Len ());return false ;};for _aef :=0;_aef < _ec .Len ();_aef ++{_bbd :=_bd .TraceToDirectObject (_ec .Get (_aef ));_cgg :=_bd .TraceToDirectObject (_bage .Get (_aef ));if _fag ,_ddc :=_bbd .(*_bd .PdfObjectDictionary );_ddc {_bc ,_ccc :=_cgg .(*_bd .PdfObjectDictionary );if !_ccc {return false ;};if !CompareDictionariesDeep (_fag ,_bc ){return false ;};}else {if _bbd .WriteString ()!=_cgg .WriteString (){_cc .Log .Debug ("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027",_bbd .WriteString (),_cgg .WriteString ());return false ;};};};continue ;};if _edd .String ()!=_daf .String (){_cc .Log .Debug ("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027",_ddg ,_edd .String (),_daf .String ());_cc .Log .Debug ("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027",_edd ,_daf );_cc .Log .Debug ("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027",_edd ,_daf );return false ;};};return true ;};func HashFile (file string )(string ,error ){_ca ,_ef :=_ab .Open (file );if _ef !=nil {return "",_ef ;};defer _ca .Close ();_db :=_c .New ();if _ ,_ef =_g .Copy (_db ,_ca );_ef !=nil {return "",_ef ;};return _gd .EncodeToString (_db .Sum (nil )),nil ;};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;};if _ ,_dd :=_b .LookPath ("\u0067\u0073");_dd !=nil {return ErrRenderNotSupported ;};return _b .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_d .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();};func CompareImages (img1 ,img2 _adc .Image )(bool ,error ){_fg :=img1 .Bounds ();_dc :=0;for _gf :=0;_gf < _fg .Size ().X ;_gf ++{for _ade :=0;_ade < _fg .Size ().Y ;_ade ++{_cg ,_fec ,_edg ,_ :=img1 .At (_gf ,_ade ).RGBA ();_cf ,_ce ,_ea ,_ :=img2 .At (_gf ,_ade ).RGBA ();if _cg !=_cf ||_fec !=_ce ||_edg !=_ea {_dc ++;};};};_bfd :=float64 (_dc )/float64 (_fg .Dx ()*_fg .Dy ());if _bfd > 0.0001{_d .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_bfd ,_dc );return false ,nil ;};return true ,nil ;};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_fb ,_ae :=HashFile (file1 );if _ae !=nil {return false ,_ae ;};_fece ,_ae :=HashFile (file2 );if _ae !=nil {return false ,_ae ;};if _fb ==_fece {return true ,nil ;};_eaa ,_ae :=ReadPNG (file1 );if _ae !=nil {return false ,_ae ;};_fgb ,_ae :=ReadPNG (file2 );if _ae !=nil {return false ,_ae ;};if _eaa .Bounds ()!=_fgb .Bounds (){return false ,nil ;};return CompareImages (_eaa ,_fgb );};func CopyFile (src ,dst string )error {_da ,_bb :=_ab .Open (src );if _bb !=nil {return _bb ;};defer _da .Close ();_ba ,_bb :=_ab .Create (dst );if _bb !=nil {return _bb ;};defer _ba .Close ();_ ,_bb =_g .Copy (_ba ,_da );return _bb ;};var (ErrRenderNotSupported =_fa .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m"););func ParseIndirectObjects (rawpdf string )(map[int64 ]_bd .PdfObject ,error ){_eg :=_bd .NewParserFromString (rawpdf );_cdf :=map[int64 ]_bd .PdfObject {};for {_cb ,_cfad :=_eg .ParseIndirectObject ();if _cfad !=nil {if _cfad ==_g .EOF {break ;};return nil ,_cfad ;};switch _bdd :=_cb .(type ){case *_bd .PdfIndirectObject :_cdf [_bdd .ObjectNumber ]=_cb ;case *_bd .PdfObjectStream :_cdf [_bdd .ObjectNumber ]=_cb ;};};for _ ,_gb :=range _cdf {_fdg (_gb ,_cdf );};return _cdf ,nil ;};func ReadPNG (file string )(_adc .Image ,error ){_ed ,_fd :=_ab .Open (file );if _fd !=nil {return nil ,_fd ;};defer _ed .Close ();return _ad .Decode (_ed );};func _fdg (_ceg _bd .PdfObject ,_dbe map[int64 ]_bd .PdfObject )error {switch _bda :=_ceg .(type ){case *_bd .PdfIndirectObject :_adce :=_bda ;_fdg (_adce .PdfObject ,_dbe );case *_bd .PdfObjectDictionary :_fdb :=_bda ;for _ ,_bddb :=range _fdb .Keys (){_bag :=_fdb .Get (_bddb );if _ceb ,_edgg :=_bag .(*_bd .PdfObjectReference );_edgg {_ee ,_afe :=_dbe [_ceb .ObjectNumber ];if !_afe {return _fa .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");};_fdb .Set (_bddb ,_ee );}else {_fdg (_bag ,_dbe );};};case *_bd .PdfObjectArray :_eeb :=_bda ;for _bad ,_dca :=range _eeb .Elements (){if _ag ,_bga :=_dca .(*_bd .PdfObjectReference );_bga {_gdc ,_eff :=_dbe [_ag .ObjectNumber ];if !_eff {return _fa .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");};_eeb .Set (_bad ,_gdc );}else {_fdg (_dca ,_dbe );};};};return nil ;};
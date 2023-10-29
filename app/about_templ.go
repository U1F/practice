// Code generated by templ@(devel) DO NOT EDIT.

package app

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func About(name string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div style=\"font-family: Arial, sans-serif; background-color: #f0f0f0; margin: 0; padding: 0;\"><div style=\"max-width: 800px; margin: 0 auto; padding: 20px; background-color: #fff; box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);\"><h1 style=\"color: #333;\">")
		if err != nil {
			return err
		}
		var_2 := `About Us`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p style=\"font-size: 16px; line-height: 1.5; color: #666;\">")
		if err != nil {
			return err
		}
		var_3 := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod laoreet quam, at fringilla lorem feugiat sit amet. Integer quis venenatis nunc. Nullam eget dolor at libero tincidunt bibendum. Vestibulum a dui ac ligula volutpat euismod. Donec ac eros a orci congue facilisis. Praesent pharetra justo sit amet nulla scelerisque, nec iaculis nisi fermentum.`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><p style=\"font-size: 16px; line-height: 1.5; color: #666;\">")
		if err != nil {
			return err
		}
		var_4 := `Nullam vel massa eget mauris bibendum venenatis. In ac nisl nec turpis tincidunt mattis. Vivamus tristique dui ut massa dictum, nec fermentum ex facilisis. Aenean auctor, urna eget malesuada consectetur, ligula nulla fermentum libero, nec sollicitudin urna lorem id turpis. Fusce volutpat, ex quis finibus pellentesque, ex risus euismod leo, et congue libero ligula vel nisi.`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

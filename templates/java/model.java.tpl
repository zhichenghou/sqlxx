{{define "model"}}
package {{.ModelPackageName}}

/**
 * Generated by sqlxx.
 */

public class {{.ModelClassName}} {
  {{- range $col := .ColumnItems}}
  private {{$col.PropType}} {{$col.PropName}};
  {{- end}}
}

{{end}}
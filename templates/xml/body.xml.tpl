{{define "body"}}

    <!-- =========================================== QUERY START =========================================== -->

    <select id="getById" resultType="{{.ModelClassName}}">
        SELECT
            {{- range $idx, $col := .ColumnItems}}
            {{- if $idx}}, {{end}} 
            {{$col.SqlName}}
            {{- end}}
        FROM
            {{.Schema}}.{{.Table}}
        WHERE id = #{id}
    </select>

    <!-- =========================================== QUERY END =========================================== -->

    <!-- =========================================== UPDATE START =========================================== -->

    <update id="update" parameterType="{{.ModelClassName}}" >
        UPDATE {{.Schema}}.{{.Table}}
        SET
            {{- range $idx, $col := .ColumnItems}}
            {{- if $idx}}, {{end}} 
            {{$col.SqlName}} = {{"#{"}}{{$.ModelParamName}}.{{$col.PropName}}}
            {{- end}}
        WHERE id = {{"#{"}}{{.ModelParamName}}.id}
    </update>


    <!-- =========================================== UPDATE END =========================================== -->

    <!-- =========================================== INSERT START =========================================== -->

    <insert id="add" parameterType="{{.ModelClassName}}" useGeneratedKeys="true" keyProperty="{{.ModelParamName}}.id">
        INSERT INTO {{.Schema}}.{{.Table}} (
            {{- range $idx, $col := .ColumnItems}}
            {{- if $idx}}, {{end}} 
            {{$col.SqlName}}
            {{- end}}
        ) VALUES (
            {{- range $idx, $col := .ColumnItems}}
            {{- if $idx}}, {{end}} 
            {{"#{"}}{{$.ModelParamName}}.{{$col.PropName}}}
            {{- end}}
        )
    </insert>

    <!-- =========================================== INSERT END =========================================== -->

    <!-- =========================================== DELETE START =========================================== -->

    <delete id="deleteById">
        DELETE 
        FROM {{.Schema}}.{{.Table}}
        WHERE id = #{id}
    </delete>

    <!-- =========================================== DELETE END =========================================== -->
{{end}}

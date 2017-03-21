{{define "mapper"}}
package {{.BasePackageName}}.infra.persistence.sql.mapper;

/**
 * Generated by sqlxx.
 */

public interface {{.ModelClassName}}Mapper {
  // =========================================== QUERY START ===========================================

  {{.ModelClassName}} getById(@Param("id") Long id);

  // =========================================== QUERY START ===========================================

  // =========================================== UPDATE START ==========================================

  Integer update(@Param("{{.ModelParamName}}") {{.ModelClassName}} {{.ModelParamName}});

  // =========================================== UPDATE END ============================================

  // =========================================== INSERT START ==========================================

  Integer add(@Param("{{.ModelParamName}}") {{.ModelClassName}} {{.ModelParamName}});

  // =========================================== INSERT END ============================================

  // =========================================== DELETE START ==========================================

  Integer deleteById(@Param("id") Long id);

  // =========================================== DELETE END ============================================
}

{{end}}
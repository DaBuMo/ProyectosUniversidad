<pxsl:sort-by-attr name="orden" order="asc">
    <pxsl:copy select="escuela/alumnos/alumno">
            <span>
                <pxsl:value-of select="alumno/nombre">
                </pxsl:value-of>
                <pxsl:value-of select="alumno/edad">
                </pxsl:value-of>
            </span>

            <div>
                <pxsl:copy-of select="/escuela/nombre">
                </pxsl:copy-of>
            </div>
    </pxsl:copy>
</pxsl:sort-by-attr>
